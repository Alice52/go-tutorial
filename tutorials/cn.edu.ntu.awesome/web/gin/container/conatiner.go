package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server2 := &http.Server{
		Addr:         ":8082",
		Handler:      router2(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 借助errgroup.Group或者自行开启两个goroutine分别启动两个服务
	g.Go(func() error {
		return server1.ListenAndServe()
	})

	g.Go(func() error {
		return server2.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func router1() http.Handler {
	e := gin.Default()
	e.GET("/ping", func(c *gin.Context) { c.String(200, "01") })

	return e
}

func router2() http.Handler {
	e := gin.Default()
	e.GET("/ping", func(c *gin.Context) { c.String(200, "02") })

	return e
}
