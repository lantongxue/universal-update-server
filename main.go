package main

import (
	"uus/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("123456"))
	r.Use(sessions.Sessions("uus", store))

	r.Static("/static", "./public")
	r.LoadHTMLGlob("./view/**")

	router := NewRouter(r)
	router.AutoRouter(&controller.IndexController{})
	router.AutoRouter(&controller.LoginController{})
	router.AutoRouter(&controller.UpdateController{})
	r.Run(":8080")
}
