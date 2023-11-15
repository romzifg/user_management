package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/romzifg/user_management/models"
	rolemodule "github.com/romzifg/user_management/modules/role.module"
	usermodule "github.com/romzifg/user_management/modules/user.module"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	models.ConnectDatabase()

	// Route
	rolemodule.Routes(router)
	usermodule.Routes(router)
	
	port := os.Getenv("PORT")
	router.Run(":" + port)
}