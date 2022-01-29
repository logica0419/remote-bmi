package router

import (
	"bytes"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/logica0419/remote-bmi/server/repository"
	"gorm.io/gorm"
)

func (r *Router) signupHandler(c echo.Context) error {
	body := new(bytes.Buffer)
	_, err := body.ReadFrom(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	name := body.String()

	existingUser, _ := r.repo.SelectUserByName(name)
	if existingUser != nil {
		return c.String(http.StatusConflict, "user already exists")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user := &repository.User{
		ID:   id,
		Name: name,
	}

	err = r.repo.InsertUser(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = id.String()
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14,
		HttpOnly: true,
	}
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (r *Router) loginHandler(c echo.Context) error {
	body := new(bytes.Buffer)
	_, err := body.ReadFrom(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	name := body.String()

	user, err := r.repo.SelectUserByName(name)
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "user not found")
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = user.ID.String()
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14,
		HttpOnly: true,
	}
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
