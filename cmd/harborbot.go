package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/harborchurchla/harborbot/internal/api"
	"github.com/harborchurchla/harborbot/internal/services"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func main() {
	log.Println("Starting flottbot")

	var wg sync.WaitGroup
	var err error

	// Only increment by one so that the app fails in the case of either a bot or API failure
	wg.Add(1)

	// Run bot
	go func() {
		defer wg.Done()
		err = runBot()
	}()

	// Run API
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
	serviceAccountJson := os.Getenv("GOOGLE_SERVICE_ACCOUNT_KEY_JSON")
	serviceAccountJson = strings.ReplaceAll(serviceAccountJson, "'", "")
	conf, err := google.JWTConfigFromJSON([]byte(serviceAccountJson), spreadsheet.Scope)
	if err != nil {
		return fmt.Errorf("error while loading google service account json: %v", err)
	}
	scheduleService := services.NewScheduleService(
		spreadsheet.NewServiceWithClient(conf.Client(context.TODO())),
		os.Getenv("SCHEDULE_SHEET_ID"),
	)

	engine := api.New(scheduleService)
	engine.POST("/slack_events/v1/events", reverseProxy("http://localhost:3000"))

	// Don't start API until bot web service starts
	err = poll(func() error {
		_, err := http.Get("http://localhost:3000")
		return err
	}, 5)
	if err != nil {
		return fmt.Errorf("flottbot web server failed to start after 5 seconds: %v", err)
	}

	return engine.Run()
}

func reverseProxy(target string) gin.HandlerFunc {
	u, err := url.Parse(target)
	if err != nil {
		log.Fatalf("error while parsing reverse proxy url: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func poll(method func() error, timeoutSeconds int) error {
	start := time.Now()
	for int(time.Now().Sub(start).Seconds()) <= timeoutSeconds {
		err := method()
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("timed out polling method after %d seconds", timeoutSeconds)
}
