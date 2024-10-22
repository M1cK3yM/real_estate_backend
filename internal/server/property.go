package server

import (
	"encoding/json"
	"net/http"
	"real_estate/internal/database"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (s *Server) getAllPropertiesHanlder(c echo.Context) error {
	properties := s.db.GetAllProperties()

	if len(properties) == 0 {
		return c.JSON(http.StatusBadRequest, "Invalid Id")
	}

	return c.JSON(http.StatusOK, properties)
}

func (s *Server) getPropertyHandler(c echo.Context) error {
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
		return c.JSON(http.StatusNotFound, "Property Not Found")
	}
	return c.JSON(http.StatusOK, property)
}

func (s *Server) addPropertyHandler(c echo.Context) error {
	propertyReq := database.Property{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&propertyReq)
	if err != nil {
		log.Print("Failed to Decode the request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error")
	}

	property, err := s.db.AddProperty(propertyReq)
	if err != nil {
		log.Print("Failed to add in database:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error ")
	}
	return c.JSON(http.StatusOK, property)
}

func (s *Server) updatePropertyHandler(c echo.Context) error {
	propertyReq := database.Property{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&propertyReq)
	if err != nil {
		log.Print("Failed to Decode the request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error")
	}

	property, err := s.db.UpdateProperty(propertyReq)
	if err != nil {
		log.Print("Failed to update in database:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error ")
	}

	return c.JSON(http.StatusOK, property)
}

func (s *Server) deletePropertyHandler(c echo.Context) error {
	id := c.QueryParam("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, "Invalid Id")
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := s.db.DeleteProperty(int64(intID)); err != nil {
		return c.JSON(http.StatusNotFound, "Property Not Found")
	}

	return c.JSON(http.StatusOK, "Property Deleted")
}
