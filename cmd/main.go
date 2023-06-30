package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
