package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (c *IndexController) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("user")
	if user == nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
