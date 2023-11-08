package shop

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {

	shop := e.Group("/shop")
	shop.GET("/goods", goodsHandler)
	shop.GET("/checkout", checkoutHandler)
}
