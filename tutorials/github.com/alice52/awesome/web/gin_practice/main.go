package main

import (
	"github.com/alice52/awesome/web/gin_practice/routers"
	"github.com/alice52/awesome/web/gin_practice/service/blog"
	"github.com/alice52/awesome/web/gin_practice/service/shop"
)

func main() {

	r := routers.Include(blog.Routers, shop.Routers)
	r.Run(":8083")
}
