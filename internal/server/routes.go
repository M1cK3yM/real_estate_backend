package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	e.GET("/properties", s.getAllPropertiesHanlder)
	e.GET("/properties/q", s.getPropertyHandler)
	e.POST("/properties", s.addPropertyHandler)
	e.PUT("/properties", s.updatePropertyHandler)
	e.DELETE("/properties", s.deletePropertyHandler)

	e.GET("/clients", s.getAllClientsHandler)
	e.GET("/clients/q", s.getClientHandler)
	e.POST("/clients", s.addClientHandler)
	e.PUT("/clients", s.updateClientHandler)
	e.DELETE("/clients", s.deleteClientHanlder)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
