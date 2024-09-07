package app

import (
	"boilerplate/model"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB      *gorm.DB
	Handler *echo.Echo
}

func New() (*App, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC", host, user, pass, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(model.Table{})
	if err != nil {
		log.Fatalln(err)
	}

	rootHandler := echo.New()

	return &App{
		DB:      db,
		Handler: rootHandler,
	}, nil
}

func (app *App) Start() error {
	return app.Handler.Start(":3000")
}
