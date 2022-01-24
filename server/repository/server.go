package repository

import (
	"net/url"

	"github.com/gofrs/uuid"
)

type Server struct {
	ID           uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	UserID       uuid.UUID `gorm:"type:char(36);not null;"`
	User         User      `gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ServerNumber int       `gorm:"type:int(1);not null"`
	Address      url.URL   `gorm:"type:varchar(2048);not null;unique"`
}

func (Server) TableName() string {
	return "server_list"
}

func (repo *Repository) SelectServersByUserID(userID uuid.UUID) ([]*Server, error) {
	var servers []*Server

	res := repo.getTx().Where("user_id = ?", userID).Find(&servers)
	if res.Error != nil {
		return nil, res.Error
	}

	return servers, nil
}

func (repo *Repository) InsertServers(servers []*Server) error {
	res := repo.getTx().Create(servers)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo *Repository) UpdateServerAddress(userID uuid.UUID, serverNumber int, address url.URL) error {
	res := repo.getTx().Model(&Server{}).Where("user_id = ? AND server_number = ?", userID, serverNumber).Update("address", address)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
