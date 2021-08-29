package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httputil"
	"net/url"
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

	r.POST("/slack_events/v1/events", ReverseProxy("http://localhost:3000"))
	return r.Run()
}

func ReverseProxy(target string) gin.HandlerFunc {
	u, err := url.Parse(target)
	if err != nil {
		log.Fatalf("error while parsing reverse proxy url: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
