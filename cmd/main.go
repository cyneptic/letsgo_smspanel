package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	controllers "github.com/cyneptic/letsgo-smspanel/controller"
	"github.com/cyneptic/letsgo-smspanel/controller/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()
	e := echo.New()
	e.Use(middleware.CustomLogger)
	controllers.AddPhoneBookRoutes(e)
	controllers.AddContactRoutes(e)
	controllers.AddWalletHRoutes(e)
	appPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := e.Start(appPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
