package main

import (
	"log"

	"github.com/cyneptic/letsgo-smspanel/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.AddSendSMSRouters(e)
	log.Fatal(e.Start(":8080"))
}
