package main

import "github.com/gin-gonic/gin"

func hello() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.String(200, "hello word")
	}
}
func main() {
	r := gin.Default()
	r.GET("/", hello())

	r.Run()
}
