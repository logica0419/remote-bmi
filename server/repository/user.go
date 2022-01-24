package repository

import "github.com/gofrs/uuid"

type User struct {
	ID   uuid.UUID `gorm:"type:char(36);not null;primaryKey" json:"id"`
	Name string    `gorm:"type:varchar(32);not null" json:"name"`
}

func (User) TableName() string {
	return "user"
}

func (repo *Repository) SelectUserByID(userID uuid.UUID) (*User, error) {
	var user *User

	res := repo.getTx().First(&user, "id = ?", userID)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) InsertUser(user *User) error {
	res := repo.getTx().Create(user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
