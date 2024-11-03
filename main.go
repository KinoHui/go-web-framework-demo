package main

import (
	"net/http"
	"web-frame-demo/jin"
)

func main() {
	r := jin.New()
	r.GET("/index", func(c *jin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *jin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello jin</h1>")
		})

		v1.GET("/hello", func(c *jin.Context) {
			// expect /hello?name=jinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *jin.Context) {
			// expect /hello/jinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *jin.Context) {
			c.JSON(http.StatusOK, jin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
