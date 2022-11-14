package user_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
	"go-flutter-bootcamp/config"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"go-flutter-bootcamp/models/failure"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func New(db *gorm.DB) Repository {
	err := db.AutoMigrate(&models.UserEntity{})
	if err != nil {
		log.Println(err.Error())
	}
	return SqlRepository{
		db: db,
	}
}

type SqlRepository struct {
	db *gorm.DB
}

func (s SqlRepository) Update(ctx context.Context, userId string, file multipart.File, req *models.UpdateProfileRequest) (*models.UserDetailResponse, error) {
	fileName := utils.UUID() + ".jpg"
	path := fmt.Sprintf("./%s/%s", config.PublicDir, fileName)
	if _, err := os.Stat(config.PublicDir); os.IsNotExist(err) {
		err := os.Mkdir(config.PublicDir, 0777)
		if err != nil {
			return nil, err
		}
	}
	fo, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	data, _ := helper.TypeConverter[models.UserEntity](&req)
	data.Photo = fmt.Sprintf("%s/%s/%s", config.BaseUrl, config.PublicDir, fileName)
	if err := s.db.WithContext(ctx).Model(models.UserEntity{}).Where("id = ?", userId).Updates(&data).Error; err != nil {
		return nil, err
	}
	var dataInDb *models.UserDetailResponse
	if err := s.db.WithContext(ctx).Raw("SELECT * FROM user WHERE id = ?", userId).Scan(&dataInDb).Error; err != nil {
		return nil, err
	}
	return dataInDb, err
}

func (s SqlRepository) Detail(ctx context.Context, userId string) (*models.UserDetailResponse, error) {
	var dataInDb *models.UserDetailResponse
	tx := s.db.WithContext(ctx).Raw("SELECT * FROM user WHERE id = ?", userId).Scan(&dataInDb)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New(failure.NotFound)
	}
	return dataInDb, nil
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
