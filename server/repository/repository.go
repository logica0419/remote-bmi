package repository

import (
	"fmt"

	"github.com/logica0419/remote-bmi/server/cmd"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func NewGormRepository(c *cmd.Config) (*Repository, error) {
	db, err := newDBConnection(c)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%w", err)
	}

	return &Repository{db: db}, nil
}

func newDBConnection(c *cmd.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.MySQL.Username, c.MySQL.Password, c.MySQL.Hostname, c.MySQL.Database) + "?parseTime=true&loc=Local&charset=utf8mb4"
	logLevel := logger.Info

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB : %w", err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	err = db.AutoMigrate(User{}, Server{}, Log{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}
