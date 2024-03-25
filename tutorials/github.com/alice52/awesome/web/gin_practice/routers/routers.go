package routers

import (
	"github.com/gin-gonic/gin"
)

type Option func(e *gin.Engine)

// Include
// 1. custom router:
//
//   - considering use []Option to store router
//   - then provider init router func
//
// 2. considering parse folder path as router prefix
func Include(ops ...Option) *gin.Engine {

	r := gin.Default()
	for _, opFunc := range ops {
		opFunc(r)
	}

	return r
}
