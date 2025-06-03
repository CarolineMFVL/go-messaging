package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string // Hashé (en prod)
}

type ThreadModel struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex"`
}
