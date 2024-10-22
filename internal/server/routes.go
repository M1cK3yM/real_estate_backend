package server

import (
	"encoding/json"
	"net/http"
	"real_estate/internal/database"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	e.GET("/properties", s.getAllPropertiesHanlder)
	e.GET("/properties/q", s.getProperty)
	e.POST("/properties", s.addPropertyHandler)
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

func (s *Server) getAllPropertiesHanlder(c echo.Context) error {
	properties := s.db.GetAllProperties()

	if len(properties) == 0 {
		return c.JSON(http.StatusOK, database.Property{})
	}

	return c.JSON(http.StatusOK, properties)
}

func (s *Server) getProperty(c echo.Context) error {
	id := c.QueryParams().Get("id")

	if id == "" {
		return c.JSON(http.StatusOK, database.Property{})
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	property, err := s.db.GetProperty(int64(intID))
	if err != nil {
		c.JSON(http.StatusNoContent, property)
	}
	return c.JSON(http.StatusOK, property)
}

func (s *Server) addPropertyHandler(c echo.Context) error {
	property := database.Property{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&property)
	if err != nil {
		log.Print("Failed to Decode the request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error")
	}
	return c.JSON(http.StatusOK, property)
}
