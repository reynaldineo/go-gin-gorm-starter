package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_CREATE_USER = "failed create user"
	MESSAGE_FAILED_LOGIN       = "login failed"
	MESSAGE_FAILED_GET_USER    = "failed get user"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "success create user"
	MESSAGE_SUCCESS_LOGIN         = "login success"
	MESSAGE_SUCCESS_GET_USER      = "succes get user"
)

var (
	ErrEmailAlreadyExist = errors.New("email already exist")
	ErrEmailNotFound     = errors.New("email not found")
	ErrPasswordNotMatch  = errors.New("password not match")
	ErrGetUserById       = errors.New("failed to get user by id")
)

type (
	UserRegisterRequest struct {
		Name       string `json:"name"`
		TelpNumber string `json:"telp_number"`
		Email      string `json:"email"`
		Password   string `json:"password"`
	}

	UserResponse struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		TelpNumber string `json:"telp_number"`
		Role       string `json:"role"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}
)
