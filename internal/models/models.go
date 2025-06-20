package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string // Hashé (en prod)
}

type ThreadModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}
