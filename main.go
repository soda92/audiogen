package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
)

var ch chan int
var is_first bool

func init() {
	ch = make(chan int)
	is_first = true
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/index.html")
	})

	files := []string{"index.html", "index.js"}
	r.LoadHTMLFiles(files...)
	for _, f := range files {
		r.GET(fmt.Sprintf("/%s", f), func(c *gin.Context) {
			c.HTML(http.StatusOK, f, gin.H{})
		})
	}

	r.GET("/play", func(c *gin.Context) {
		if is_first {
			is_first = false
			ch <- 1
		} else {
			ch <- 0
			ch <- 1
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	return r
}

func main() {
	r := setupRouter()
	addr := "127.0.0.1:8080"
	go browser.OpenURL(fmt.Sprintf("http://%s", addr))

	go sdl2_play(ch)
	// Listen and Serve
	r.Run(addr)
}
