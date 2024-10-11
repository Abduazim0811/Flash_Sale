package user


type User struct {
    ID       string  `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Price    float32 `json:"price"`
}

type UserRequest struct {
    Username        string `json:"username"`
    Password        string `json:"password"`
    ConfirmPassword string `json:"confirm_password"`
    Email           string `json:"email"`
}

type UserResponse struct {
    ID       string  `json:"id"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
}

type GetUserRequest struct {
    ID string `json:"id"`
}

type UserEmpty struct{}

type ListUser struct {
    Users []User `json:"users"`
}

type UpdateUserReq struct {
    ID       string  `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Price    float32 `json:"price"`
}

type UpdateUserRes struct {
    Message string `json:"message"`
}

type UpdatePasswordReq struct {
    ID          string  `json:"id"`
    OldPassword string `json:"old_password"`
    NewPassword string `json:"new_password"`
}

type VerifyReq struct {
    Email string `json:"email"`
    Code  string `json:"code"`
}
