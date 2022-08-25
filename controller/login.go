package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (c *LoginController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
