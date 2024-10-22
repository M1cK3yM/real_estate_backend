package server

import (
	"encoding/json"
	"log"
	"net/http"
	"real_estate/internal/database"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) getAllClientsHandler(c echo.Context) error {
	clients := s.db.GetAllClients()

	if len(clients) == 0 {
		return c.JSON(http.StatusBadRequest, "Invalid Id")
	}

	return c.JSON(http.StatusOK, clients)
}

func (s *Server) getClientHandler(c echo.Context) error {
	id := c.QueryParams().Get("id")

	if id == "" {
		return c.JSON(http.StatusOK, database.Client{})
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	client, err := s.db.GetClient(int64(intID))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Client Not Found")
	}
	return c.JSON(http.StatusOK, client)
}

func (s *Server) addClientHandler(c echo.Context) error {
	clientReq := database.Client{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&clientReq)
	if err != nil {
		log.Print("Failed to Decode the request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error")
	}

	client, err := s.db.AddClient(clientReq)
	if err != nil {
		log.Print("Failed to add in database:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error ")
	}
	return c.JSON(http.StatusOK, client)
}

func (s *Server) updateClientHandler(c echo.Context) error {
	clientReq := database.Client{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&clientReq)
	if err != nil {
		log.Print("Failed to Decode the request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error")
	}

	client, err := s.db.UpdateClient(clientReq)
	if err != nil {
		log.Print("Failed to update in database:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An Unexpected Error ")
	}

	return c.JSON(http.StatusOK, client)
}

func (s *Server) deleteClientHanlder(c echo.Context) error {
	id := c.QueryParam("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, "Invalid Id")
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := s.db.DeleteClient(int64(intID)); err != nil {
		return c.JSON(http.StatusNotFound, "Client Not Found")
	}

	return c.JSON(http.StatusOK, "Client Deleted")
}
