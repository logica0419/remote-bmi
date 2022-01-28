package router

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/logica0419/remote-bmi/server/benchmark"
	"github.com/logica0419/remote-bmi/server/repository"
	"github.com/sapphi-red/go-traq"
	"github.com/srinathgs/mysqlstore"
)

type Router struct {
	e        *echo.Echo
	address  string
	cli      *traq.APIClient
	clientID string
	repo     *repository.Repository
	bench    *benchmark.Benchmarker
}

type Config struct {
	Address  string
	Version  string
	ClientID string
}

func NewRouter(cfg *Config, repo *repository.Repository, bench *benchmark.Benchmarker, db *sql.DB) (*Router, error) {
	e, err := newEcho(db)
	if err != nil {
		return nil, err
	}

	cli := traq.NewAPIClient(traq.NewConfiguration())

	r := &Router{
		e:        e,
		address:  cfg.Address,
		cli:      cli,
		clientID: cfg.ClientID,
		repo:     repo,
		bench:    bench,
	}

	api := r.e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
		api.GET("/version", func(c echo.Context) error {
			return c.String(http.StatusOK, cfg.Version)
		}, checkLoginMiddleware)

		oauth := api.Group("/oauth")
		{
			oauth.GET("/callback", r.getOauthCallbackHandler)
			oauth.POST("/code", r.postOAuthCodeHandler)
		}

		user := api.Group("/users", checkLoginMiddleware)
		{
			user.GET("/me", r.getUsersMeHandler)
		}

		server := api.Group("/servers", checkLoginMiddleware)
		{
			server.GET("", r.getServersHandler)
			server.POST("", r.postServersHandler)
			server.PUT("/:server_number", r.putServersServerNumberHandler)
			server.DELETE("", r.deleteServersHandler)
		}

		api.POST("/benchmark", r.postBenchmarkHandler, checkLoginMiddleware)
	}

	r.e.File("/oauth", "client/dist/index.html")
	r.e.Static("/", "client/dist")

	return r, nil
}

func newEcho(db *sql.DB) (*echo.Echo, error) {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${prefix} ${short_file} ${line} |")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339} method = ${method} | uri = ${uri} | status = ${status} ${error}\n"}))

	store, err := mysqlstore.NewMySQLStoreFromConnection(db, "session", "/", 3600, []byte("<SecretKey>"))
	if err != nil {
		return nil, err
	}
	e.Use(session.Middleware(store))

	return e, nil
}

func (r *Router) Run() {
	r.e.Logger.Fatal(r.e.Start(r.address))
}
