package main

import (
	"fmt"
	"log"
	"os"
	controller "quest/controller/api"
	"quest/repository"
	"quest/routes"
	"quest/service"

	model "quest/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// base_url: "169.51.207.94:31290/"

func main() {

	envErr := godotenv.Load("config.env")
	if envErr != nil {
		log.Fatal(envErr)
	}

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DBHost"),
		os.Getenv("DBUser"),
		os.Getenv("DBPassword"),
		os.Getenv("DBName"),
		os.Getenv("DBPort"),
	)

	gormDb, err := gorm.Open(postgres.New(postgres.Config{DSN: connStr}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	gormDb.AutoMigrate(model.Citizen{})
	gormDb.AutoMigrate(model.Document{})
	gormDb.AutoMigrate(model.Operator{})

	citizenRepository := repository.NewCitizenRepository(gormDb)
	citizenService := service.NewCitizenService(citizenRepository)
	citizenController := controller.NewCitizenController(citizenService)

	documentRepository := repository.NewDocumentRepository(gormDb)
	documentService := service.NewDocumentService(documentRepository)
	documentController := controller.NewDocumentController(documentService)

	operatorRepository := repository.NewOperatorRepository(gormDb)
	operatorService := service.NewOperatorService(operatorRepository)
	operatorController := controller.NewOperatorController(operatorService)

	app := gin.Default()

	routes.RegisterCitizenRoutes(app, citizenController)
	routes.RegisterDocumentRoutes(app, documentController)
	routes.RegisterOperatorRoutes(app, operatorController)

	app.Run(":3000")
}
