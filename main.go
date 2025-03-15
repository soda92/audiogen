package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
)

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
		go play_sound()
		c.JSON(http.StatusOK, gin.H{})
	})


	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	addr := "127.0.0.1:8080"
	go browser.OpenURL(fmt.Sprintf("http://%s", addr))
	r.Run(addr)
}
