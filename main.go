package main

import (
	"github.com/gin-gonic/gin"
	"sms-api/controller"
)

func main() {
	r := gin.Default()
	r.GET("/", controller.Index)
	r.POST("/sms", controller.PostSMS)
	r.Run()
}
