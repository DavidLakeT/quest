package main

import (
	controller "quest/controller/api"
	"quest/repository"
	"quest/routes"
	"quest/service"

	model "quest/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// base_url: "169.51.207.94:31290/"

func main() {

	gormDb, err := gorm.Open(sqlite.Open("../database/carpetaciudadana.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	gormDb.AutoMigrate(model.Citizen{})
	gormDb.AutoMigrate(model.Document{})
	gormDb.AutoMigrate(model.Operator{})

	citizenRepository := repository.NewCitizenRepository(gormDb)
	citizenService := service.NewCitizenService(citizenRepository)
	citizenController := controller.NewCitizenController(citizenService)

	app := gin.Default()

	routes.RegisterCitizenRoutes(app, citizenController)
	app.Run(":3000")
}
