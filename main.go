package main

import (
	"uus/config"

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

	middleware := NewMiddleware(r)
	middleware.Register(config.Middlewares)

	router := NewRouter(r)
	router.Register(config.Routers)
	r.Run(":8080")
}
