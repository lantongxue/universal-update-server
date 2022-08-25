package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (c *LoginController) Index(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		account := ctx.PostForm("account")
		if account == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"error":   http.StatusInternalServerError,
				"message": "Account cannot be empty",
			})
			return
		}
		password := ctx.PostForm("password")
		if password == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"error":   http.StatusInternalServerError,
				"message": "Password cannot be empty",
			})
			return
		}
		if account != "admin" && password != "admin" {
			ctx.JSON(http.StatusOK, gin.H{
				"error":   http.StatusInternalServerError,
				"message": "Account or Password error",
			})
			return
		}
		session := sessions.Default(ctx)
		session.Set("user", "5555")
		session.Save()
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}

	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
