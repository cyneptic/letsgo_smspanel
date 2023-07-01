package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func CustomLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		stop := time.Now()
		latency := stop.Sub(start)

		if err != nil {
			c.Error(err)
		} else {
			req := c.Request()
			res := c.Response()

			fmt.Printf("[%v] %s(%s) \"%s %s\" %d %d (%v)\n", stop.Format("2006-01-02 15:04:05"), c.RealIP(), c.Path(), req.Method, req.URL.String(), res.Status, res.Size, latency)
		}

		return err
	}
}
