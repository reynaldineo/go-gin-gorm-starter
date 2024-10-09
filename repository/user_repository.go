package repository

import (
	"context"

	"github.com/reynaldineo/go-gin-gorm-starter/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetUserByEmail(ctx context.Context, email string) (entity.User, bool, error)
		CreateUser(ctx context.Context, user entity.User) (entity.User, error)
		GetUserById(ctx context.Context, userId string) (entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, bool, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}
	return user, true, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, userId string) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("id = ?", userId).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
