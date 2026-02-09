package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	model "golan-quickstart/models"
	userrepo "golan-quickstart/repository"
	service "golan-quickstart/server"
	utils "golan-quickstart/utils"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserInfoRoutes(r *gin.RouterGroup) {
	userinfo := r.Group("/userinfo")
	{
		userinfo.POST("/login", LoginUser)
		userinfo.POST("/signup", SignupUser)
	}
}

func SignupUser(c *gin.Context) {
	var users model.UserInformation

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	savedUser, err := service.UserVerification(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(savedUser.Id)
	fmt.Println("token" + token + "token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "User signup successful",
		"jwt_token": token,
		"user": gin.H{
			"id":    savedUser.Id,
			"name":  savedUser.Name,
			"email": savedUser.Email,
		},
	})
}

func LoginUser(c *gin.Context) {
	var req Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}
	fmt.Println("Login attempt:", req)
	// Fetch user
	db_user, err := userrepo.GetUserByEmail(req.Email)
	if err != nil {
		fmt.Println("User not found error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User is not found"})
		return
	}
	fmt.Println("User found:", db_user)
	fmt.Println("Stored password hash:", db_user.Password)
	if !utils.CheckPasswordHash(req.Password, db_user.Password) {
		fmt.Println("Password hash mismatch")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(db_user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"name":  db_user.Name,
			"email": db_user.Email,
		},
	})
}
