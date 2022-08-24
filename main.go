package main

import (
	"general-update-server/controller"
	"general-update-server/controller/a/b"
	"general-update-server/controller/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router := NewRouter(r)
	router.AutoRouter(&controller.UpdateController{})
	router.AutoRouter(&user.IndexController{})
	router.AutoRouter(&b.IndexController{})
	r.Run(":8080")
}
