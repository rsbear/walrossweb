package main

import (
	"context"
	"embed"
	"fmt"
	"io"
	"net/http"
	"strings"
	"walross/walrossweb/pkgs/web"

	"github.com/charmbracelet/log"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static/*
var static embed.FS

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Server{
		echo: e,
	}
}

func (s *Server) serveStaticFile(c echo.Context) error {
	path := strings.TrimPrefix(c.Request().URL.Path, "/static/")
	path = "static/" + path
	log.Debug("serving static file", "path", path)

	content, err := static.ReadFile(path)
	if err != nil {
		log.Error("failed to read static file", "error", err)
		return echo.NotFoundHandler(c)
	}

	if strings.HasSuffix(path, ".css") {
		c.Response().Header().Set("Content-Type", "text/css; charset=utf-8")
	}

	if strings.HasSuffix(path, ".js") {
		c.Response().Header().Set("Content-Type", "text/javascript; charset=utf-8")
	}

	if strings.HasSuffix(path, ".ico") {
		c.Response().Header().Set("Content-Type", "image/x-icon")
	}

	_, err = io.WriteString(c.Response(), string(content))
	if err != nil {
		log.Error("failed to write response", "error", err)
	}
	return err
}

func (s *Server) setupRoutes() {
	s.echo.GET("/static/*", s.serveStaticFile)
	s.echo.GET("/", s.handleLandingRoute)

	// /<namespace> and /<namespace>?readme<repo>
	// s.echo.GET("/:namespace", s.handleNamespaceRoute)
	// s.echo.GET("/:namespace/readme", s.handleNamespaceReadmeRoute)

}

func (s *Server) handleLandingRoute(c echo.Context) error {
	page := web.BaseTemplate(LandingRoute())
	return s.HTML(c, page)
}

func LandingRoute() elem.Node {
	return elem.Div(attrs.Props{attrs.Class: "text-center"},
		elem.H1(attrs.Props{attrs.Class: "text-4xl"}, elem.Text("Walross")),
	)
}

func (s *Server) Start() error {
	s.setupRoutes()
	addr := fmt.Sprintf("%s:%d", "localhost", 8080)
	log.Debug("server starting", "addr", addr)
	return s.echo.Start(addr)
}

func (s *Server) HTML(c echo.Context, el elem.Node) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return c.HTML(http.StatusOK, el.Render())
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.echo != nil {
		return s.echo.Shutdown(ctx)
	}
	return nil
}

func main() {
	s := NewServer()

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("server started")
}
