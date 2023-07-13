package main

import (
	"log"

	controllers "github.com/cyneptic/letsgo-smspanel/controller"
	"github.com/cyneptic/letsgo-smspanel/controller/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.CustomLogger)
	controllers.AddPhoneBookRoutes(e)
	controllers.AddSendSMSRouters(e)
	controllers.AddContactRoutes(e)
	controllers.AddTemplateRoutes(e)
	controllers.AddWalletHRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
