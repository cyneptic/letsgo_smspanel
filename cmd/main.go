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
	_ = godotenv.Load(".env")
	e := echo.New()
	e.Use(middleware.CustomLogger)
	controllers.AddPhoneBookRoutes(e)
	controllers.AddSendSMSRouters(e)
	controllers.AddContactRoutes(e)
	controllers.AddTemplateRoutes(e)
	controllers.AddWalletHRoutes(e)
	controllers.RegisterNumberHandler(e)
  controllers.AddAdminActionRoutes(e)
	appPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := e.Start(appPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
