package routers

import (
	_ "email_code/docs"
	"email_code/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/ping", services.Ping)
	r.POST("/send-code", services.SendCode)
	r.POST("/register", services.Register)
	
	r.Run(":9090")
}
