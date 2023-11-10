package main

import (
	"github.com/alice52/awesome/web/gin_practice/docs"
	"github.com/alice52/awesome/web/gin_practice/routers"
	"github.com/alice52/awesome/web/gin_practice/service/blog"
	"github.com/alice52/awesome/web/gin_practice/service/shop"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://github.com/alice52

// @contact.name alice52
// @contact.url https://github.com/alice52
// @contact.email https://github.com/alice52

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8083
// @BasePath
func main() {

	r := routers.Include(blog.Routers, shop.Routers)
	_ = docs.SwaggerInfo.BasePath

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8083")
}
