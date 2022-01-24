package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

type Config struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

func NewRepository(c *Config) (*Repository, error) {
	err := createDB(c)
	if err != nil {
		return nil, err
	}

	db, err := newDBConnection(c)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%d", err)
	}

	return &Repository{db: db}, nil
}

func createDB(c *Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", c.Username, c.Password, c.Hostname, c.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + c.Database)
	if err != nil {
		return err
	}

	return nil
}

func newDBConnection(c *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Hostname, c.Port, c.Database) + "?parseTime=true&loc=Local&charset=utf8mb4"
	logLevel := logger.Info

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB: %d", err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	err = db.AutoMigrate(User{}, Server{}, Log{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %d", err)
	}

	return db, nil
}

func GetSqlDB(repo *Repository) (*sql.DB, error) {
	return repo.db.DB()
}

func (repo *Repository) getTx() *gorm.DB {
	return repo.db.Session(&gorm.Session{})
}
