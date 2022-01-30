package router

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/logica0419/remote-bmi/server/repository"
)

type serversResponse struct {
	ID           uuid.UUID `json:"id,omitempty"`
	ServerNumber int       `json:"server_number,omitempty"`
	Address      string    `json:"address,omitempty"`
}

func (r *Router) getServersHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	servers, err := r.repo.SelectServersByUserID(userUUID)
	if len(servers) == 0 {
		return c.String(http.StatusNotFound, "servers not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := []serversResponse{}
	for _, v := range servers {
		res = append(res, serversResponse{
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
	if len(req) > 3 {
		return c.String(http.StatusBadRequest, "too many servers")
	}

	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	existingServers, _ := r.repo.SelectServersByUserID(userUUID)
	if len(existingServers) > 0 {
		return c.String(http.StatusConflict, "servers already exist")
	}

	servers := []*repository.Server{}
	serverNumberUsed := map[int]bool{}
	for _, server := range req {
		id, err := uuid.NewV4()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		verifyAddress := server.Address
		if !strings.HasPrefix(server.Address, "http://") && !strings.HasPrefix(server.Address, "https://") {
			verifyAddress = "http://" + server.Address
		}
		_, err = url.Parse(verifyAddress)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if server.ServerNumber < 1 || server.ServerNumber > 3 {
			return c.String(http.StatusBadRequest, "server number must be 1, 2 or 3")
		}

		if _, ok := serverNumberUsed[server.ServerNumber]; ok {
			return c.String(http.StatusBadRequest, "server_number is duplicated")
		}
		serverNumberUsed[server.ServerNumber] = true

		servers = append(servers, &repository.Server{
			ID:           id,
			UserID:       userUUID,
			ServerNumber: server.ServerNumber,
			Address:      server.Address,
		})
	}

	if err := r.repo.InsertServers(servers); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := []serversResponse{}
	for _, v := range servers {
		res = append(res, serversResponse{
			ID:           v.ID,
			ServerNumber: v.ServerNumber,
			Address:      v.Address,
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (r *Router) putServersServerNumberHandler(c echo.Context) error {
	body := new(bytes.Buffer)
	_, err := body.ReadFrom(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	address := body.String()

	verifyAddress := address
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		verifyAddress = "http://" + address
	}
	_, err = url.Parse(verifyAddress)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	serverNumberStr := c.Param("server_number")
	serverNumber, err := strconv.Atoi(serverNumberStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if serverNumber < 1 || serverNumber > 3 {
		return c.String(http.StatusBadRequest, "server number must be 1, 2 or 3")
	}

	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	err = r.repo.UpdateServerAddress(userUUID, serverNumber, address)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (r *Router) deleteServersHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userUUID, _ := uuid.FromString(sess.Values["user_id"].(string))

	existingServers, _ := r.repo.SelectServersByUserID(userUUID)
	if len(existingServers) == 0 {
		return c.String(http.StatusNotFound, "servers not found")
	}

	err := r.repo.DeleteServersByUserID(userUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
