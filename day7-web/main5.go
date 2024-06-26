package main

import (
	"log"
	"net/http"
	"time"

	"github.com/davidhong1/go-playground/day7-web/gee"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main5() {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware
	r.GET("/index", func(c *gee.Context) {
		c.HTMLRaw(http.StatusOK, "<h1>Index Page</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
