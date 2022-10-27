package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 2. 注册中间件
	r.Use(costStat())

	r.Run(":8084")
}

// 1. 定义中间件
func costStat() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Set("env", "enc")
		c.Next() // proceed || c.Abort()

		cost := time.Since(start)
		log.Println(cost)
	}
}
