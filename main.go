package main

import (
	"os"
	"os/exec"
	"sync"
)
import "log"

func main() {
	log.Println("Starting flottbott")

	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command("./flottbot")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
	}()

	wg.Wait()
	if err != nil {
		log.Fatalf("flotbott failed: %v", err)
	}
}
