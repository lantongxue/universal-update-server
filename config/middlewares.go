package config

import (
	"net/http"
	"strings"
	"uus/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Middlewares map[string]gin.HandlerFunc = map[string]gin.HandlerFunc{
	"auth": auth,
}

func auth(ctx *gin.Context) {
	if strings.HasPrefix(ctx.Request.RequestURI, "/login") {
		ctx.Next()
		return
	}
	session := sessions.Default(ctx)
	loginToken := session.Get("login-token").(model.LoginToken)
	if !loginToken.IsValidate() {
		session.Clear()
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
