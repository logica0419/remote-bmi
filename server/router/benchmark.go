package router

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type postBenchmarkRequest struct {
	ServerNumber int `json:"server_number,omitempty"`
}

func (r *Router) postBenchmarkHandler(c echo.Context) error {
	req := postBenchmarkRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	id, err := r.bench.Run(userUUID, req.ServerNumber)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log, err := r.repo.SelectLogByID(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := getLogsResponse{
		ID:           log.ID,
		ServerNumber: log.Server.ServerNumber,
		StdOut:       log.StdOut,
		CreatedAt:    log.CreatedAt,
	}

	return c.JSON(http.StatusOK, res)
}
