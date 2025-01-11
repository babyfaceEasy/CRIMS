package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/xo/dburl"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var c *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	var err error

	dbUrl, err := dburl.Parse(fmt.Sprintf("%s?sslmode=%s&timezone=%s", os.Getenv("DATABASE_URL"), os.Getenv("DB_SSL_MODE"), os.Getenv("TIMEZONE")))
	if err != nil {
		return nil, err
	}

	c, err = gorm.Open(postgres.Open(dbUrl.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := c.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 20)

	return c, nil
}

func RunMigrations() error {
	db, err := c.DB()
	if err != nil {
		return err
	}
	log.Print("Running migrations")

	if err := goose.Up(db, "db/migrations"); err != nil {
		return err
	}
	return nil
}
