package server

import (
	"os"

	"github.com/babyfaceeasy/crims/internal/db"
	"github.com/babyfaceeasy/crims/internal/repository"
	"github.com/babyfaceeasy/crims/internal/routes"
	"github.com/babyfaceeasy/crims/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var engine *gin.Engine

// Initialize setup the server with the needed details
func Initialize() error {
	if err := LoadEnv(); err != nil {
		return err
	}

	// connect to database
	DB, err := db.ConnectToDatabase()
	if err != nil {
		return err
	}

	// run migrations
	if err := db.RunMigrations(); err != nil {
		return err
	}

	engine = gin.Default()

	repo := repository.NewRepository(DB)
	svc := services.NewService(repo)

	// Register routes
	routes.RegisterRoutes(engine, svc)

	return nil
}

func LoadEnv() error {
	return godotenv.Load(".env")
}

// Run the app
func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	engine.Run(":" + port)
}
