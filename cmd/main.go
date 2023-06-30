package main

import (
	"log"

	controllers "github.com/cyneptic/letsgo-smspanel/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controllers.AddPhoneBookRoutes(e)
	controllers.AddContactRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
