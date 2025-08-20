package api

import (
	"fmt"
	"go_admin_api/result"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context) {
	fmt.Println(c)
	result.Success(c, 200)
}

func Failed(c *gin.Context) {
	result.Failed(c, int(result.ApiCode.Failed), result.ApiCode.GetMessage(result.ApiCode.Failed))
}
