package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateController struct{}

func (u *UpdateController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "update/index")
}

// bad method
func (u *UpdateController) Index2() {
	fmt.Println("asdasdad")
}

func (u *UpdateController) Post__PUT(ctx *gin.Context) {

}
