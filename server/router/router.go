package router

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func SetupEcho() *echo.Echo {
	e := newEcho()

	api := e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(200, "pong")
		})
		api.GET("/version", func(c echo.Context) error {
			return c.String(200, "0.0.1")
		})
	}

	e.Static("/", "client/dist")

	return e
}

func newEcho() *echo.Echo {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${prefix} ${short_file} ${line} |")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339} method = ${method} | uri = ${uri} | status = ${status} ${error}\n"}))

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	return e
}
