package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (c *LoginController) Index(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		username := ctx.PostForm("username")
		if username == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Username cannot be empty",
			})
			return
		}
		password := ctx.PostForm("password")
		if password == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Password cannot be empty",
			})
			return
		}
		if username != "admin" && password != "admin" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Username or Password error",
			})
			return
		}
		session := sessions.Default(ctx)
		session.Set("user", "5555")
		session.Save()
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
			"url":     "/",
		})
		return
	}
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
