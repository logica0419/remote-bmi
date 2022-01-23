package repository

import (
	"net/url"

	"github.com/gofrs/uuid"
)

type Server struct {
	ID           uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	UserID       uuid.UUID `gorm:"type:char(36);not null"`
	User         User      `gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ServerNumber int       `gorm:"type:int(1);not null"`
	Address      url.URL   `gorm:"type:varchar(2048);not null"`
}

func (Server) TableName() string {
	return "server_list"
}
