package main

import (
	"net/http"
	"web-frame-demo/jin"
)

func main() {
	r := jin.New()
	r.GET("/", func(c *jin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Jin</h1>")
	})
	r.GET("/hello", func(c *jin.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *jin.Context) {
		c.JSON(http.StatusOK, jin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
