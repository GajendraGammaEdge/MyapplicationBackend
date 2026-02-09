package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	controllers "golan-quickstart/controllers"
	Db "golan-quickstart/dbconfig"
	auth "golan-quickstart/middlerware"
	utils "golan-quickstart/utils"
)

func main() {
	Db.ConnectDatabase()
	if err := utils.LoadRSAKeys(); err != nil {
		fmt.Println("Failed to load RSA keys:", err)
		return
	}
	fmt.Println("Golang Gin Framework is Working Successfully")
	gin.ForceConsoleColor()
	router := gin.Default()

	// Serve frontend static files at /frontend
	router.Static("/frontend", "./frontend")

	controllers.UserInfoRoutes(router.Group("/gogin"))
	protected := router.Group("/pro")

	protected.Use(auth.AuthMiddleware())
	{
		controllers.UserInformation(protected)
		controllers.HealthCheckRoutes(protected)
	}

	router.Run(":8080")
}
