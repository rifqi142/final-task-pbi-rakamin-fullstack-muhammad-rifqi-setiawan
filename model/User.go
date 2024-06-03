package model

import "time"

type User struct {
	ID        uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string  `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email     string  `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required"`
	Password  string  `gorm:"not null" json:"password" form:"password" valid:"required~Password is required"`
	Photos    []Photo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE" json:"photos" form:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Custom Request Response Request for User
type CreateUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	Username string `json:"username" form:"username" binding:"required"`
}
type UpdateUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Username string `json:"username" form:"username" binding:"required"`
}
type LoginUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

// Custom Response for User
type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserResponse struct {
	ID 	  		uint     `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type UserPhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
