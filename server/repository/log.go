package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type Log struct {
	ID        uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	UserID    uuid.UUID `gorm:"type:char(36);not null"`
	User      User      `gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ServerID  uuid.UUID `gorm:"type:char(36);not null"`
	Server    Server    `gorm:"foreignkey:ServerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StdOut    string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
}

func (Log) TableName() string {
	return "benchmark_log"
}
