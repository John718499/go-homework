package main

import (
	rh "homework/internal/RoleHelper"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/role", rh.GetAllRole)

	router.GET("/role/:id", rh.GetOneRole)

	router.POST("/role", rh.AddRole)

	router.PUT("/role/:id", rh.UpdateRole)

	router.DELETE("/role/:id", rh.DeleteRole)

	router.Run(":8080")
}
