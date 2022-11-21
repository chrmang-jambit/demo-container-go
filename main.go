package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chrmang-jambit/demo-container-go/pkg/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//go:generate oapi-codegen --config oapi-config.yaml api/openapi.yaml

const (
	cfg_serverip   = "SERVER_IP"
	cfg_serverport = "SERVER_PORT"
	path_health    = "/health"
	path_ready     = "/ready"
)

type config struct {
	serverip   string
	serverport int
}

func main() {
	maincfg := readConfig()
	mainServer := echo.New()
	mainServer.HideBanner = true

	corrcfg := middleware.RequestIDConfig{}
	mainServer.Use(middleware.RequestIDWithConfig(corrcfg))

	logcfg := middleware.LoggerConfig{
		Skipper: urlSkipper,
	}
	mainServer.Use(middleware.LoggerWithConfig(logcfg))
	mainServer.Logger.SetLevel(log.DEBUG)

	mainServerUri := fmt.Sprintf("%s:%d", maincfg.serverip, maincfg.serverport)

	mainServer.GET(path_health, healthcheck)
	mainServer.GET(path_ready, readycheck)

	server := api.New()
	api.RegisterHandlers(mainServer, server)
	mainServer.Logger.Fatal(mainServer.Start(mainServerUri))
}

func healthcheck(ctx echo.Context) error {
	ctx.Response().Status = http.StatusOK
	return nil
}

func readycheck(ctx echo.Context) error {
	ctx.Response().Status = http.StatusOK
	return nil
}

func readConfig() config {
	viper.SetDefault(cfg_serverip, "localhost")
	viper.SetDefault(cfg_serverport, 8080)

	viper.AutomaticEnv()

	return config{
		serverip:   viper.GetString(cfg_serverip),
		serverport: viper.GetInt(cfg_serverport),
	}
}

func urlSkipper(c echo.Context) bool {
	// ignore health endpoint for logging and metrics
	return strings.HasPrefix(c.Path(), path_health) || strings.HasPrefix(c.Path(), path_ready)
}
