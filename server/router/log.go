package router

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type getLogsResponse struct {
	ID           uuid.UUID `json:"id,omitempty"`
	ServerNumber int       `json:"server_number,omitempty"`
	StdOut       string    `json:"stdout,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

func (r *Router) getLogsHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	logs, err := r.repo.SelectLogsByUserID(userUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := []*getLogsResponse{}
	for _, log := range logs {
		res = append(res, &getLogsResponse{
			ID:           log.ID,
			ServerNumber: log.Server.ServerNumber,
			StdOut:       log.StdOut,
			CreatedAt:    log.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, res)
}
