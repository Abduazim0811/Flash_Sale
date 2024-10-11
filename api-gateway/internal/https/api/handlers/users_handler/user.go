package usershandler

import (
	"api-gateway/internal/protos/user_proto"
	"context"
	"net/http"
	"time"

	_ "api-gateway/docs"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	ClientUser user_proto.UserServiceClient
}

// @title Flash-Sale
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @securityDefinitions.apikey Bearer
// @in 				header
// @name Authorization
// @description Enter the token in the format `Bearer {token}`
// @host localhost:9876
// @BasePath /

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body user_proto.UserRequest true "User request body"
// @Success 200 {object} user_proto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /register [post]
func (u *UsersHandler) Register(c *gin.Context) {
	var req user_proto.UserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.Register(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// VerifyCode godoc
// @Summary Verify a user code
// @Description Verify the user code sent to the user's email
// @Tags user
// @Accept json
// @Produce json
// @Param verify body user_proto.VerifyReq true "Verification request body"
// @Success 200 {object} user_proto.UserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /verify-code [post]
func (u *UsersHandler) VerifyCode(c *gin.Context) {
	var req user_proto.VerifyReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.VerifyCode(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary User login
// @Description Log in a user with email and password
// @Tags user
// @Accept json
// @Produce json
// @Param login body user_proto.LoginRequest true "Login request body"
// @Success 200 {object} user_proto.LoginResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /login [post]
func (u *UsersHandler) Login(c *gin.Context) {
	var req user_proto.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetbyIdUsers godoc
// @Summary Get user by ID
// @Description Get a user by their unique ID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user_proto.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id} [get]
func (u *UsersHandler) GetbyIdUsers(c *gin.Context) {
	userId := c.Param("id")
	var req user_proto.GetUserRequest
	req.Id = userId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.GetByIdUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all registered users
// @Tags user
// @Produce json
// @Success 200 {object} user_proto.ListUser
// @Failure 500 {object} string
// @Security Bearer
// @Router /users [get]
func (u *UsersHandler) GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.GetUsers(ctx, &user_proto.UserEmpty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUsers godoc
// @Summary Update a user
// @Description Update a user's details by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param update body user_proto.UpdateUserReq true "Update request body"
// @Success 200 {object} user_proto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id} [put]
func (u *UsersHandler) UpdateUsers(c *gin.Context) {
	userId := c.Param("id")
	var req user_proto.UpdateUserReq
	req.Id = userId
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.UpdateUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUsersPassword godoc
// @Summary Update user password
// @Description Update a user's password by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param update body user_proto.UpdatePasswordReq true "Password update request body"
// @Success 200 {object} user_proto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id}/password [put]
func (u *UsersHandler) UpdateUsersPassword(c *gin.Context) {
	userId := c.Param("id")
	var req user_proto.UpdatePasswordReq
	req.Id = userId

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.UpdatePassword(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteUsers godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user_proto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id} [delete]
func (u *UsersHandler) DeleteUsers(c *gin.Context){
	userId := c.Param("id")
	var req user_proto.GetUserRequest
	req.Id = userId
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.DeleteUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}