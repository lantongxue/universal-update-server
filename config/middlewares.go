package config

import "github.com/gin-gonic/gin"

var Middlewares map[string]gin.HandlerFunc = map[string]gin.HandlerFunc{
	"auth": auth,
}

func auth(ctx *gin.Context) {

}
