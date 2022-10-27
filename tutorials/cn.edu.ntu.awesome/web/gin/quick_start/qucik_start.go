package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 2. request
	// 2.1 simple
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	// 2.2 complex
	router.POST("/login", func(c *gin.Context) {
		var login Login
		// ShouldBind() 会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 3. 重定向
	// 3.1 HTTP
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})
	// 3.2 路由
	router.GET("/test", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/ping"
		router.HandleContext(c)
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8083")
}
