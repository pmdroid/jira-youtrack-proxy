package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/pmdroid/jira-youtrack-proxy/internal/config"
	"github.com/pmdroid/jira-youtrack-proxy/internal/handler"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Failed to build config")
		
		panic(1)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/rest/api/2/issue", func(c echo.Context) error {
		return handler.HandleCreateIssue(c, cfg)
	})

	e.GET("/rest/api/2/project/:id", func(c echo.Context) error {
		return handler.HandleProjectDetails(c, cfg)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	fmt.Printf("Starting proxy server on port %s\n", cfg.Port)
	fmt.Printf("YouTrack URL: %s\n", cfg.YouTrackURL)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
