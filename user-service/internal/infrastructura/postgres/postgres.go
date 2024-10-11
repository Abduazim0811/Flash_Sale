package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"user-service/internal/entity/user"
	"user-service/internal/infrastructura/repository"
	"user-service/protos/notification_proto"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserPostgres struct {
	db *sql.DB
	N notification_proto.NotificationServiceClient
}

func NewUserPostgres(db *sql.DB, N notification_proto.NotificationServiceClient) repository.UserRepository {
	return &UserPostgres{db: db, N: N}
}

func (u *UserPostgres) InsertUsersPostgres(req user.UserRequest) (*user.UserResponse, error) {
	userid := uuid.New().String()
	var res user.UserResponse
	sql, args, err := squirrel.
		Insert("users").
		Columns("id, username", "email", "password").
		Values(userid, req.Username, req.Email, req.Password).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error generating SQL for AddUser: %v", err)
	}

	_, err = u.N.SendNotification(context.Background(),&notification_proto.SendNotificationRequest{UserId: userid, Message: fmt.Sprintf("your account is created successfully with this is %v you can get notifications with it", userid)})
	if err != nil{
		log.Println(err)
	}
	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&res.ID); err != nil {
		return nil, fmt.Errorf("error scanning result in AddUser: %v", err)
	}
	return &res, nil
}

func (u *UserPostgres) GetbyEmailUsersPostgres(email string) (*user.User, error) {
	var res user.User
	sql, args, err := squirrel.
		Select("*").
		From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("email not found: ", err)
		return nil, fmt.Errorf("email not found: %v", err)
	}

	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&res.ID, &res.Username, &res.Email, &res.Password, &res.Price); err != nil {
		log.Println("scan error")
		return nil, fmt.Errorf("scan error: %v", err)
	}

	return &res, nil

}

func (u *UserPostgres) GetbyIdUserPostgres(req user.GetUserRequest) (*user.User, error) {
	var res user.User
	sqls, args, err := squirrel.
		Select("*").
		From("users").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error generating SQL for GetByIdUser: %v", err)
	}

	row := u.db.QueryRow(sqls, args...)
	if err := row.Scan(&res.ID, &res.Username, &res.Email, &res.Password, &res.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found for GetByIdUser: %v", err)
		}
		return nil, fmt.Errorf("error scanning result in GetByIdUser: %v", err)
	}

	return &res, nil
}

func (u *UserPostgres) GetAllUsersPostgres() (*user.ListUser, error) {
	sql, args, err := squirrel.
		Select("*").
		From("users").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error generating SQL for GetAll: %v", err)
	}

	rows, err := u.db.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query for GetAll: %v", err)
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Price); err != nil {
			return nil, fmt.Errorf("error scanning row in GetAll: %v", err)
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows in GetAll: %v", err)
	}

	return &user.ListUser{Users: users}, nil
}

func (u *UserPostgres) UpdateUser(req user.UpdateUserReq) error {
	updateBuilder := squirrel.Update("users")

	var isUpdated bool 

	if req.Username != "" {
		updateBuilder = updateBuilder.Set("username", req.Username)
		isUpdated = true 
	}

	if req.Email != "" {
		updateBuilder = updateBuilder.Set("email", req.Email)
		isUpdated = true 
	}

	if req.Price != 0 {
		updateBuilder = updateBuilder.Set("price", req.Price)
		isUpdated = true 
	}
	fmt.Println(req.Price)
	if !isUpdated {
		return fmt.Errorf("no fields to update")
	}

	updateBuilder = updateBuilder.Where(squirrel.Eq{"id": req.ID})

	sql, args, err := updateBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for UpdateUser: %v", err)
	}

	result, err := u.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing SQL in UpdateUser: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows in UpdateUser: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated, user with ID %v might not exist", req.ID)
	}

	_, err = u.N.SendNotification(context.Background(), &notification_proto.SendNotificationRequest{
		UserId:  req.ID,
		Message: fmt.Sprintf("your account has been updated successfully with ID %v", req.ID),
	})
	if err != nil {
		log.Println(err)
	}

	return nil
}



func (u *UserPostgres) UpdatePassword(req user.UpdatePasswordReq) error {
	var currentPassword string
	sqlSelect, argsSelect, err := squirrel.
		Select("password").
		From("users").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for password retrieval: %v", err)
	}

	err = u.db.QueryRow(sqlSelect, argsSelect...).Scan(&currentPassword)
	if err != nil {
		return fmt.Errorf("error retrieving current password: %v", err)
	}
	fmt.Println("old password:", req.OldPassword)
	err = bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(req.OldPassword))
	if err != nil {
		return fmt.Errorf("error comparing old password: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing new password: %v", err)
	}

	sqlUpdate, argsUpdate, err := squirrel.
		Update("users").
		Set("password", string(hashedPassword)).
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for UpdatePassword: %v", err)
	}

	_, err = u.db.Exec(sqlUpdate, argsUpdate...)
	if err != nil {
		return fmt.Errorf("error executing SQL in UpdatePassword: %v", err)
	}

	return nil
}

func (u *UserPostgres) DeleteUser(req user.GetUserRequest) error {
	var exists bool
	err := u.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", req.ID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking user existence: %v", err)
	}

	if !exists {
		return fmt.Errorf("user with ID %s does not exist", req.ID)
	}

	sql, args, err := squirrel.
		Delete("users").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for DeleteUser: %v", err)
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing SQL in DeleteUser: %w", err)
	}

	return nil
}
