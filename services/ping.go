package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ping测试",
	})
}