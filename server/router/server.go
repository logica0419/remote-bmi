package router

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/logica0419/remote-bmi/server/repository"
)

type getServersResponse struct {
	ID           uuid.UUID `json:"id,omitempty"`
	ServerNumber int       `json:"server_number,omitempty"`
	Address      url.URL   `json:"address,omitempty"`
}

func (r *Router) getServersHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	servers, err := r.repo.SelectServersByUserID(userUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := []getServersResponse{}
	for _, v := range servers {
		res = append(res, getServersResponse{
			ID:           v.ID,
			ServerNumber: v.ServerNumber,
			Address:      v.Address,
		})
	}

	return c.JSON(http.StatusOK, res)
}

type postServersRequest struct {
	ServerNumber int    `json:"server_number,omitempty"`
	Address      string `json:"address,omitempty"`
}

func (r *Router) postServersHandler(c echo.Context) error {
	req := []postServersRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	servers := []*repository.Server{}
	serverNumberUsedMap := map[int]bool{}
	for _, v := range req {
		id, err := uuid.NewV4()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		address, err := url.Parse(v.Address)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if _, ok := serverNumberUsedMap[v.ServerNumber]; ok {
			return c.String(http.StatusBadRequest, "server_number is duplicated")
		}
		serverNumberUsedMap[v.ServerNumber] = true

		servers = append(servers, &repository.Server{
			ID:           id,
			UserID:       userUUID,
			ServerNumber: v.ServerNumber,
			Address:      *address,
		})
	}

	if err := r.repo.InsertServers(servers); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, servers)
}

func (r *Router) putServersServerNumberHandler(c echo.Context) error {
	body := new(bytes.Buffer)
	_, err := body.ReadFrom(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	addressStr := body.String()
	address, err := url.Parse(addressStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	serverNumberStr := c.Param("server_number")
	serverNumber, err := strconv.Atoi(serverNumberStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	err = r.repo.UpdateServerAddress(userUUID, serverNumber, *address)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
