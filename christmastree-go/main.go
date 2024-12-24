package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	port := 3000

	go func() {
		fmt.Printf("Server started at http://localhost:%d\n", port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	go func() {
		draw(4)
	}()

	waitForShutdown()
}

func draw(max int) {
	if max < 1 {
		fmt.Println("The number of levels must be at least 1.")
		return
	}

	totalWidth := (max * 2) - 1
	center := totalWidth / 2

	for i := 0; i < max; i++ {
		stars := strings.Repeat("*", (i*2)+1)
		row := pad(stars, center+i+1, totalWidth)
		fmt.Println(row)
	}

	verticalLine := pad("|", center+1, totalWidth)
	fmt.Println(verticalLine)
}

func pad(content string, leftPadding int, totalWidth int) string {
	padded := fmt.Sprintf("%*s", leftPadding, content)
	return fmt.Sprintf("%-*s", totalWidth, padded)
}

func waitForShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	sig := <-signalChan
	fmt.Printf("Received signal: %v. Shutting down gracefully...\n", sig)
}
