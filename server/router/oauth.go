package router

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/logica0419/remote-bmi/server/repository"
	"github.com/sapphi-red/go-traq"
	"github.com/thanhpk/randstr"
)

const callbackBaseURL = "https://q.trap.jp/api/v3/oauth2/authorize"

func (r *Router) getOauthCallbackHandler(c echo.Context) error {
	verifier := randstr.String(64)
	hash := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(hash[:])

	sess, _ := session.Get("session", c)
	sess.Values["verifier"] = verifier
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   30,
		HttpOnly: true,
	}
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	uri := fmt.Sprintf("%s?response_type=code&client_id=%s&code_challenge=%s&code_challenge_method=S256", callbackBaseURL, r.clientID, challenge)
	return c.String(http.StatusOK, uri)
}

func (r *Router) postOAuthCodeHandler(c echo.Context) error {
	body := new(bytes.Buffer)
	_, err := body.ReadFrom(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	code := body.String()

	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	verifier := sess.Values["verifier"]
	if verifier == nil {
		return c.String(http.StatusBadRequest, "verifier is not found")
	}
	verifierStr := verifier.(string)
	opts := &traq.Oauth2ApiPostOAuth2TokenOpts{
		Code:         optional.NewString(code),
		ClientId:     optional.NewString(r.clientID),
		CodeVerifier: optional.NewString(verifierStr),
	}
	token, res, err := r.cli.Oauth2Api.PostOAuth2Token(context.Background(), "authorization_code", opts)
	if err != nil || token.AccessToken == "" || res.StatusCode >= 400 {
		return c.String(res.StatusCode, err.Error())
	}

	auth := context.WithValue(context.Background(), traq.ContextAccessToken, token.AccessToken)
	me, res, err := r.cli.MeApi.GetMe(auth)
	if err != nil || me.Id == "" || me.Name == "" || res.StatusCode >= 400 {
		return c.String(res.StatusCode, err.Error())
	}

	meUUID, _ := uuid.FromString(me.Id)
	user, _ := r.repo.SelectUserByID(meUUID)
	if user == nil {
		user := &repository.User{
			ID:   meUUID,
			Name: me.Name,
		}
		err := r.repo.InsertUser(user)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	sess = sessions.NewSession(sess.Store(), "session")
	sess.Values["user_id"] = me.Id
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(token.ExpiresIn),
		HttpOnly: true,
	}
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
