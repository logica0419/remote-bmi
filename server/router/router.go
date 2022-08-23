package router

import (
	"crypto/rand"
	"database/sql"
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/logica0419/remote-bmi/server/benchmark"
	"github.com/logica0419/remote-bmi/server/repository"
	"github.com/srinathgs/mysqlstore"
)

//go:embed static
var embeddedFiles embed.FS

func getFileSystem() http.FileSystem {
	fSys, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(fSys)
}

type Router struct {
	e       *echo.Echo
	address string
	repo    *repository.Repository
	bench   *benchmark.Benchmarker
}

type Config struct {
	Address string
	Version string
}

func NewRouter(cfg *Config, repo *repository.Repository, bench *benchmark.Benchmarker, db *sql.DB) (*Router, error) {
	e, err := newEcho(db)
	if err != nil {
		return nil, err
	}

	r := &Router{
		e:       e,
		address: cfg.Address,
		repo:    repo,
		bench:   bench,
	}

	api := r.e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
		api.GET("/version", func(c echo.Context) error {
			return c.String(http.StatusOK, cfg.Version)
		}, checkLoginMiddleware)

		auth := api.Group("")
		{
			auth.POST("/signup", r.signupHandler)
			auth.POST("/login", r.loginHandler)
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

		api.GET("/logs", r.getLogsHandler, checkLoginMiddleware)
	}

	assetHandler := http.FileServer(getFileSystem())
	r.e.GET("/*", echo.WrapHandler(assetHandler))

	return r, nil
}

func newEcho(db *sql.DB) (*echo.Echo, error) {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${prefix} ${short_file} ${line} |")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339} method = ${method} | uri = ${uri} | status = ${status} ${error}\n"}))

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	var secKey string
	for _, v := range b {
		secKey += string(letters[int(v)%len(letters)])
	}

	store, err := mysqlstore.NewMySQLStoreFromConnection(db, "session", "/", 3600, []byte(secKey))
	if err != nil {
		return nil, err
	}
	e.Use(session.Middleware(store))

	return e, nil
}

func (r *Router) Run() {
	r.e.Logger.Fatal(r.e.Start(r.address))
}
