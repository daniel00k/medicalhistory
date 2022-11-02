package main

import (
	"github.com/daniel00k/medicalhistory/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	uc := controller.NewUsersController()
	router := gin.Default()
	router.POST("/users", uc.Create)
	router.GET("/users", uc.Index)
	router.Run()
}
