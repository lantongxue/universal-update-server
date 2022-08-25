package main

import (
	"uus/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./public")
	r.LoadHTMLGlob("./view/**")
	router := NewRouter(r)
	router.AutoRouter(&controller.UpdateController{})
	router.AutoRouter(&controller.LoginController{})
	r.Run(":8080")
}
