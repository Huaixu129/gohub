package main

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 处理404
	r.NoRoute(func(c *gin.Context) {
		// 获取标头的 Accept 信息
		accept := c.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			// 如果是 html 的话
			c.String(http.StatusNotFound, "404 页面")
		} else {
			// 默认返回 json
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义, 请确认 url 和请求方法是否正确",
			})
		}
	})

	r.Run(":8080")
}