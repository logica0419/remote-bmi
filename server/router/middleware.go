package router

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func checkLoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		userID, ok := sess.Values["user_id"].(string)
		if !ok || userID == "" {
			return c.String(http.StatusUnauthorized, "please login")
		}
		_, err := uuid.FromString(userID)
		if err != nil {
			return c.String(http.StatusUnauthorized, "invalid userID")
		}

		return next(c)
	}
}
