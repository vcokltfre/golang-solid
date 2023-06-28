package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vcokltfre/golang-solid/src/frontend"
)

func Start(bind string) error {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(frontend.WebMiddleware)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "HTTP[${time_rfc3339}] ${status} ${method} ${uri} ${latency_human} ${error}\n",
	}))

	return e.Start(bind)
}
