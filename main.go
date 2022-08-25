package main

import (
	"uus/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./public")
	router := NewRouter(r)
	router.AutoRouter(&controller.UpdateController{})
	r.Run(":8080")
}
