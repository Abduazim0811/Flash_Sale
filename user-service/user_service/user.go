package userservice

import (
	"context"
	"fmt"
	"log"
	"user-service/internal/service"
	"user-service/protos/user_proto"
)

type UsersGrpc struct {
	user_proto.UnimplementedUserServiceServer
	s service.UserService
}

func NewUsersGrpc(s service.UserService) *UsersGrpc {
	return &UsersGrpc{s: s}
}

func (u *UsersGrpc) Register(ctx context.Context, req *user_proto.UserRequest) (*user_proto.UpdateUserRes, error) {
	res, err := u.s.Registeruser(req)
	if err != nil {
		log.Println("register users error:", err)
		return nil, fmt.Errorf("register users error: %v", err)
	}

	return res, nil
}

func (u *UsersGrpc) Login(ctx context.Context, req *user_proto.LoginRequest) (*user_proto.LoginResponse, error) {
	res, err := u.s.LoginUser(req)
	if err != nil {
		log.Println("login error: ", err)
		return nil, fmt.Errorf("login error: %v", err)
	}

	return res, nil
}

func (u *UsersGrpc) VerifyCode(ctx context.Context, req *user_proto.VerifyReq) (*user_proto.UserResponse, error) {
	res, err := u.s.Verifycode(req)
	if err != nil {
		log.Println("verify error: ", err)
		return nil, fmt.Errorf("verify error: %v", err)
	}

	return res, nil
}

func (u *UsersGrpc) GetByIdUser(ctx context.Context, req *user_proto.GetUserRequest) (*user_proto.User, error) {
	res, err := u.s.GetByIDusers(req)
	fmt.Println(res.Password, res.Price)
	if err != nil {
		log.Println("get by id error: ", err)
		return nil, fmt.Errorf("get by id error: %v", err)
	}

	return res, nil
}

func (u *UsersGrpc) GetUsers(ctx context.Context, req *user_proto.UserEmpty) (*user_proto.ListUser, error) {
	res, err := u.s.GetAllusers(req)
	if err != nil {
		log.Println("get all users:", err)
		return nil, fmt.Errorf("get all users: %v", err)
	}

	return res, nil
}

func (u *UsersGrpc) UpdateUser(ctx context.Context, req *user_proto.UpdateUserReq) (*user_proto.UpdateUserRes, error) {
	res, err := u.s.Updateuser(req)
	if err != nil {
		log.Println("update users error: ", err)
		return nil, fmt.Errorf("update users error: %v", err)
	}
	return res, nil
}

func (u *UsersGrpc) UpdatePassword(ctx context.Context, req *user_proto.UpdatePasswordReq) (*user_proto.UpdateUserRes, error) {
	res, err := u.s.UpdatePassword(req)
	if err != nil {
		log.Println("update password error:", err)
		return nil, fmt.Errorf("update password error: %v", err)
	}

	return res, nil
}

func (u *UsersGrpc) DeleteUser(ctx context.Context, req *user_proto.GetUserRequest) (*user_proto.UpdateUserRes, error) {
	res, err := u.s.Deleteusers(req)
	if err != nil {
		log.Println("delete users error: ", err)
		return nil, fmt.Errorf("delete users error: %v", err)
	}

	return res, nil
}
