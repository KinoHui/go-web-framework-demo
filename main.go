package main

import (
	"net/http"
	"web-frame-demo/jin"
)

func main() {
	r := jin.New()
	r.GET("/", func(c *jin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *jin.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *jin.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *jin.Context) {
		c.JSON(http.StatusOK, jin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
