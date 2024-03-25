package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary do api
// @Schemes
// @Description do post
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /post [get]
func postHandler(c *gin.Context) {

	c.JSON(http.StatusOK, "post")
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "post")
}
