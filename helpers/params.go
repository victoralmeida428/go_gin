package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryInt(ctx *gin.Context, param string) (int, error){
	vString := ctx.Query(param)
	return strconv.Atoi(vString)
}