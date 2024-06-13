package main

import (
	"net/http"

	"github.com/davidhong1/go-playground/day7-web/gee"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[1000])
	})

	r.Run(":9999")
}
