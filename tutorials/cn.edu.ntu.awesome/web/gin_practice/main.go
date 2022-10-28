package main

import (
	"cn.edu.ntu.awesome/web/gin_practice/routers"
	"cn.edu.ntu.awesome/web/gin_practice/service/blog"
	"cn.edu.ntu.awesome/web/gin_practice/service/shop"
)

func main() {

	r := routers.Include(blog.Routers, shop.Routers)
	r.Run(":8083")
}
