package service

import (
	"context"
	"sync"

	"github.com/reynaldineo/go-gin-gorm-starter/constant"
	"github.com/reynaldineo/go-gin-gorm-starter/dto"
	"github.com/reynaldineo/go-gin-gorm-starter/entity"
	"github.com/reynaldineo/go-gin-gorm-starter/helper"
	"github.com/reynaldineo/go-gin-gorm-starter/repository"
	"gorm.io/gorm"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserRegisterRequest) (dto.UserResponse, error)
		VerifyUser(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		GetUserByID(ctx context.Context, userId string) (dto.UserResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

var (
	mu sync.Mutex
)

func (s *userService) RegisterUser(ctx context.Context, req dto.UserRegisterRequest) (dto.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, isEmailExist, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.UserResponse{}, err
	}

	if isEmailExist {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExist
	}

	user := entity.User{
		Name:       req.Name,
		Email:      req.Email,
		TelpNumber: req.TelpNumber,
		Password:   req.Password,
		Role:       constant.ENUM_ROLE_USER,
	}

	userRes, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:         userRes.ID.String(),
		Name:       userRes.Name,
		Email:      userRes.Email,
		TelpNumber: userRes.TelpNumber,
		Role:       userRes.Role,
	}, nil
}

func (s *userService) VerifyUser(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, isEmailExist, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil || !isEmailExist {
		return dto.UserLoginResponse{}, dto.ErrEmailNotFound
	}

	isPasswordMatch, err := helper.CompareHashPassword(user.Password, []byte(req.Password))
	if err != nil || !isPasswordMatch {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	token := s.jwtService.GenerateToken(user.ID.String(), user.Role)

	return dto.UserLoginResponse{
		Token: token,
		Role:  user.Role,
	}, nil

}

func (s *userService) GetUserByID(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		ID:         user.ID.String(),
		Name:       user.Name,
		Email:      user.Email,
		TelpNumber: user.TelpNumber,
		Role:       user.Role,
	}, nil
}
