package repository

import "user-service/internal/entity/user"

type UserRepository interface {
	InsertUsersPostgres(req user.UserRequest) (*user.UserResponse, error)
	GetbyEmailUsersPostgres(email string) (*user.User, error)
	GetbyIdUserPostgres(req user.GetUserRequest) (*user.User, error)
	GetAllUsersPostgres() (*user.ListUser, error)
	UpdateUser(req user.UpdateUserReq) error
	UpdatePassword(req user.UpdatePasswordReq) error
	DeleteUser(req user.GetUserRequest) error
}
