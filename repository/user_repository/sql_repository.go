package user_repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/models/failure"
	"gorm.io/gorm"
)

func New(db *gorm.DB) Repository {
	db.AutoMigrate(&models.UserEntity{})
	return SqlRepository{
		db: db,
	}
}

type SqlRepository struct {
	db *gorm.DB
}

func (s SqlRepository) Detail(ctx context.Context, userId string) (*models.UserDetailResponse, error) {
	var dataInDb *models.UserEntity
	tx := s.db.WithContext(ctx).Raw("SELECT * FROM user WHERE id = ?", userId).Scan(&dataInDb)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New(failure.NotFound)
	}
	result, _ := helper.TypeConverter[models.UserDetailResponse](dataInDb)
	return result, nil
}

func (s SqlRepository) Create(ctx context.Context, req *models.CreateUserRequest) error {
	var userInDb *models.UserEntity
	tx := s.db.WithContext(ctx).Raw("SELECT * FROM user WHERE email = ?", req.Email).Scan(&userInDb)
	if tx.RowsAffected > 0 {
		return errors.New(failure.AlreadyExists)
	}
	data, _ := helper.TypeConverter[models.UserEntity](&req)
	password, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}
	data.Id = uuid.NewString()
	data.Password = password
	data.Photo = "https://ui-avatars.com/api/?name=" + req.Name + "&background=EBF4FF&size=128"
	tx = s.db.WithContext(ctx).Model(models.UserEntity{}).Create(&data)
	return tx.Error
}

func (s SqlRepository) Login(ctx context.Context, req *models.LoginRequest) (*models.UserDetailResponse, error) {
	var userInDb *models.UserEntity
	tx := s.db.WithContext(ctx).Raw("SELECT * FROM user WHERE email = ?", req.Email).Scan(&userInDb)
	if tx.RowsAffected == 0 {
		return nil, errors.New(failure.InvalidCredential)
	}
	err := helper.CheckPassword(req.Password, userInDb.Password)
	if err != nil {
		return nil, errors.New(failure.InvalidCredential)
	}
	result, _ := helper.TypeConverter[models.UserDetailResponse](&userInDb)
	return result, nil
}
