package router

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (r *Router) getUsersMeHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	user, err := r.repo.SelectUserByID(userUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
