package repository

import "github.com/gofrs/uuid"

type User struct {
	ID   uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Name string    `gorm:"type:varchar(32);not null"`
}

func (User) TableName() string {
	return "user"
}
