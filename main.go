package main

import (
	"net/http"
	"web-frame-demo/jin"
)

func main() {
	r := jin.Default()
	r.GET("/", func(c *jin.Context) {
		c.String(http.StatusOK, "Hello jin\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *jin.Context) {
		names := []string{"jin"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
