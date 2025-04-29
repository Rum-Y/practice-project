package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userURL, _ := url.Parse("http://user-service:8080")
	r.Any("/user/*path", createProxy(userURL))

	productURL, _ := url.Parse("http://product-service:8081")
	r.Any("/product/*path", createProxy(productURL))

	r.Run(":8000")
}

func createProxy(target *url.URL) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return func(c *gin.Context) {
		c.Request.URL.Path = c.Param("path")
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
