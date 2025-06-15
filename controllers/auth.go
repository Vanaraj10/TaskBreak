package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Vanaraj10/taskmorph-backend/config"
	"github.com/Vanaraj10/taskmorph-backend/models"
	"github.com/Vanaraj10/taskmorph-backend/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	userCollection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existing models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}
	// Hash the password here
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var reqUser models.User
	
	c.BindJSON(&reqUser)
	userCollection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dbUser models.User
	err := userCollection.FindOne(ctx,bson.M{"email": reqUser.Email}).Decode(&dbUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(reqUser.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	token, _ := services.GenerateToken(dbUser.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		 "token":   token,
	})
}