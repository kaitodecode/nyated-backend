package dto

type UserResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Role  string  `json:"role"`
	Token *string `json:"token,omitempty"`
}

type UserRegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
	RoleID          string
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UpdateUserRequest struct {
	Name            string `json:"name" validate:"required,min=3"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password,omitempty" validate:"min=6"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"min=6"`
	OldPassword     string `json:"old_password,omitempty" validate:"min=6"`
	RoleID          string
}
