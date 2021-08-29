package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"sync"
)

func main() {
	log.Println("Starting flottbot")

	var wg sync.WaitGroup
	var err error

	// Run bot
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = runBot()
	}()

	// Run API
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = runApi()
	}()

	wg.Wait()
	if err != nil {
		log.Fatalf("bot failed: %v", err)
	}
}

func runBot() error {
	cmd := exec.Command("./flottbot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runApi() error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/slack_events/v1/events", ReverseProxyToFlottbot("localhost:3000"))
	return r.Run()
}

func ReverseProxyToFlottbot(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
