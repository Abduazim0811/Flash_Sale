package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"user-service/internal/entity/user"
	"user-service/internal/infrastructura/redis"
	"user-service/internal/infrastructura/repository"
	pkg "user-service/internal/pkg/email"
	"user-service/internal/pkg/jwt"
	"user-service/protos/user_proto"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo  repository.UserRepository
	redis *redis.RedisClient
}

func NewUserService(repo repository.UserRepository, redis *redis.RedisClient) *UserService {
	return &UserService{repo: repo, redis: redis}
}

func (u *UserService) Registeruser(req *user_proto.UserRequest) (*user_proto.UpdateUserRes, error) {
	var userreq user.UserRequest
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("password and confirm password do not match")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	req.Password = string(bytes)
	code := 10000 + rand.Intn(90000)
	err = pkg.SendEmail(req.Email, pkg.SendClientCode(code, req.Username))
	if err != nil {
		return nil, fmt.Errorf("error sending email to user: %v", err)
	}

	userreq.Username = req.Username
	userreq.Email = req.Email
	userreq.Password = req.Password

	userData := map[string]interface{}{
		"userName": userreq.Username,
		"email":    userreq.Email,
		"password": userreq.Password,
		"code":     code,
	}

	if u.redis == nil {
		log.Println("redis error", err)
		return nil, fmt.Errorf("redis client is not initialized")
	}

	err = u.redis.SetHash(req.Email, userData)
	if err != nil {
		return nil, fmt.Errorf("failed to save user data in Redis: %v", err)
	}

	return &user_proto.UpdateUserRes{Message: "Verify code"}, nil
}

func (u *UserService) Verifycode(req *user_proto.VerifyReq) (*user_proto.UserResponse, error) {
	res, err := u.redis.VerifyEmail(context.Background(), req.Email, req.Code)
	if err != nil {
		log.Println("verify code error: ")
		return nil, fmt.Errorf("verify code error: %v", err)
	}

	var userreq user.UserRequest

	userreq.Username = res.Username
	userreq.Email = res.Email
	userreq.Password = res.Password

	userres, err := u.repo.InsertUsersPostgres(userreq)
	if err != nil {
		log.Println("error")
		return nil, fmt.Errorf("error: %v", err)
	}

	return &user_proto.UserResponse{
		Id: userres.ID,
	}, nil
}

func (u *UserService) LoginUser(req *user_proto.LoginRequest) (*user_proto.LoginResponse, error) {
	res, err := u.repo.GetbyEmailUsersPostgres(req.Email)
	if err != nil {
		log.Println("login error")
		return nil, fmt.Errorf("login erro")
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	token, err := jwt.GenerateJWTToken(req.Email)
	if err != nil {
		return nil, err
	}

	return &user_proto.LoginResponse{Token: token}, nil
}

func (u *UserService) GetByIDusers(req *user_proto.GetUserRequest)(*user_proto.User, error){
	userData, err := u.redis.GetHash(req.Id)
    if err == nil && len(userData) > 0 {
		fmt.Println(userData)
        return &user_proto.User{
            Id:       req.Id,
            Username: userData["username"],
            Email:    userData["email"],
            Password: userData["password"], 
        }, nil
    }
	res, err := u.repo.GetbyIdUserPostgres(user.GetUserRequest{ID: req.Id})
	if err != nil{
		log.Println("get by id users postgres error: ", err)
		return nil, fmt.Errorf("get by id users postgres error: %v", err)
	}

	return &user_proto.User{
		Id: res.ID,
		Username: res.Username,
		Email: res.Email,
		Password: res.Password,
	}, nil
}

func (u *UserService) GetAllusers(req *user_proto.UserEmpty)(*user_proto.ListUser, error){
	res, err := u.repo.GetAllUsersPostgres()
	if err != nil {
		return nil, fmt.Errorf("error fetching all users: %v", err)
	}
	var protoUsers []*user_proto.User
	for _, u := range res.Users {
		protoUser := &user_proto.User{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
			Price: u.Price,
		}
		protoUsers = append(protoUsers, protoUser)
	}

	return &user_proto.ListUser{User: protoUsers}, nil
}

func (u *UserService) Updateuser(req *user_proto.UpdateUserReq)(*user_proto.UpdateUserRes, error){
	err := u.repo.UpdateUser(user.UpdateUserReq{
		ID: req.Id,
		Username: req.Username,
		Email: req.Email,
		Price: req.Price,
	})
    if err != nil {
		log.Println("error updating user in datavase: ", err)
        return nil, fmt.Errorf("error updating user in database: %v", err)
    }

    userDataToCache := map[string]interface{}{
        "username": req.Username,
        "email":    req.Email,
    }

    err = u.redis.SetHash(req.Id, userDataToCache)
    if err != nil {
        log.Println("failed to update cache in Redis:", err)
        return nil, fmt.Errorf("failed to update cache in Redis: %v", err)
    }

    return &user_proto.UpdateUserRes{Message: "User updated successfully"}, nil
}

func (u *UserService) UpdatePassword(req *user_proto.UpdatePasswordReq)(*user_proto.UpdateUserRes, error){
	err := u.repo.UpdatePassword(user.UpdatePasswordReq{
		ID: req.Id,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		log.Println("update password users error:", err)
		return nil, fmt.Errorf("update password users error: %v", err)
	}
	userDataToCache := map[string]interface{}{
        "password": req.NewPassword,
    }
	err = u.redis.SetHash(req.Id, userDataToCache)
    if err != nil {
        log.Println("failed to update cache in Redis:", err)
        return nil, fmt.Errorf("failed to update cache in Redis: %v", err)
    }

    return &user_proto.UpdateUserRes{Message: "User updated successfully"}, nil
}

func (u *UserService) Deleteusers(req *user_proto.GetUserRequest)(*user_proto.UpdateUserRes, error){
	err := u.repo.DeleteUser(user.GetUserRequest{ID: req.Id})
	if err != nil {
		log.Println("delete users error: ", err)
		return nil, fmt.Errorf("delete users error: %v",err)
	}

	err = u.redis.Delete(req.Id)
	if err != nil {
		log.Println("delete redis error:", err)
		return nil, fmt.Errorf("delete redis error: %v", err)
	}

	return &user_proto.UpdateUserRes{Message: "users deleted"}, nil
}