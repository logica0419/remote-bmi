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

func (repo *Repository) SelectLogsByUserID(userID uuid.UUID) ([]*Log, error) {
	var logs []*Log

	res := repo.getTx().Joins("User").Joins("Log").Where("user_id = ?", userID).Find(&logs)
	if res.Error != nil {
		return nil, res.Error
	}

	return logs, nil
}

func (repo *Repository) InsertLog(log Log) error {
	res := repo.getTx().Create(&log)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
