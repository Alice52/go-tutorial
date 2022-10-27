## gin quick-start

1. 简介

   - 基于 httprouter 开发的 Web 框架
   - 文档健全, 性能优异

2. install

   ```go
   go get -u github.com/gin-gonic/gin
   ```

3. core point

   - 路由: 普通路由 | 路由组
   - 中间件(拦截器): 定义 | 注册 | 使用
     1. gin 默认使用了 Logger & Recovery
     2. 为某个路由单独注册
     3. 为路由组注册中间件
     4. 注册全局中间件
   - 容器: 结果可以是 json | xml | yaml |protobuf | html 等

4. 请求参数: `c.ShouldBind(&login)`

   ```go
   // 1. 获取querystring参数: /user/search?username=xx
   username := c.DefaultQuery("username", "小王子")

   // 2. 获取form参数: /user/search
   username := c.PostForm("username")

   // 3. 获取json参数, 之后再反序列化
   b, _ := c.GetRawData()

   // 4. 获取path参数: /user/search/小王子/
   r.GET("/user/search/:username", func(c *gin.Context) {
         username := c.Param("username")
   }
   ```
