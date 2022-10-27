package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. suport restful api: GET | DELETE | PUT | POST
	router := gin.Default()

	// 2. 普通路由
	router.Any("/ping", func(ctx *gin.Context) {})

	// 3. 路由组: 支持嵌套
	sysGroup := router.Group("/sys")
	{
		sysGroup.POST("/check", func(c *gin.Context) {})
		userGroup := router.Group("/user")
		{
			userGroup.POST("/check", func(c *gin.Context) {})
		}
	}

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8083")
}
