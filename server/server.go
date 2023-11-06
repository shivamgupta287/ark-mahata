package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"otp/go-gin-poc/database"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("SERVER_ADDRESS")

	database.ConnectDB()

	router := gin.Default()

	initializeRouter(router)

	logrus.Error(router.Run(port))

}
