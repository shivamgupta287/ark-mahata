package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"net/http"
	"otp/go-gin-poc/database"
	"otp/go-gin-poc/models"
	"time"
)

func SingleLoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		filter := bson.M{"login_id": user.LoginId}
		err := database.UserColl.FindOne(context.TODO(), filter)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"your OTP is": generateOTP(),
			})
		}

	}
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(10000)
	return fmt.Sprintf("%04d", otp)
}
