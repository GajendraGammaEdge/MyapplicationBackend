package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "golan-quickstart/server"
)

func UserInformation(r *gin.RouterGroup) {
	userinfo := r.Group("/userinfo")
	{
		userinfo.GET("/:id", GetUserInfoByID)
		userinfo.GET("/getall", GetUserAll)
	}
}

func GetUserInfoByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is not found",
		})
		return
	}
	ls := service.GetUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get api is work successfully",
		"user":    ls,
	})
}

func GetUserAll(c *gin.Context) {
	lst := service.GetAllUser()
	c.JSON(http.StatusOK, gin.H{
		"users": lst,
	})
}
