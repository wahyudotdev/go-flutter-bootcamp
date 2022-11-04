package models

type UserEntity struct {
	Id       string `json:"id" gorm:"primaryKey,column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (r UserEntity) TableName() string {
	return "user"
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required" form:"name"`
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required,min=8" form:"password"`
}
