package infraestructure

import (
	//"errors"
	"github.com/fatih/color"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

type (
	Server struct {
		Postgres Postgres
	}
	Welcome struct {
		Message string
	}
)

func (s *Server) Run() {
	e := echo.New()

	e.Get("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Welcome{"Welcome to Go-Sample-Rest by DaemonGear"})
	}

	color.Green("Server Run ------------------- http://localhost:9090")
	e.Run(standard.New(":9090"))

}
