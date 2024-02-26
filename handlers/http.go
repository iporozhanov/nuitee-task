package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	app    APP
	Router *gin.Engine
}

// APP is an interface for the route handlers
type APP interface {
	GetHotels(*GetHotelsRequest) (*GetHotelsResponse, error)
}

func NewHTTP(app APP) *HTTP {
	return &HTTP{
		app: app,
	}
}

// InitRoutes initializes the routes for the Gin HTTP server
func (h *HTTP) InitRoutes() {
	r := gin.Default()

	r.GET("/hotels", h.GetHotels)

	h.Router = r
}

// Run starts the HTTP server
func (h *HTTP) Run(port string) {
	h.Router.Run(fmt.Sprintf(":%s", port))
}

// GetHotelsRequest is the request for the GetHotels route
func (h *HTTP) GetHotels(c *gin.Context) {
	var req GetHotelsRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	// Check if the required query parameters are present
	if c.Query("hotelIds") == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{"hotelIds is required"})
		return
	}
	if c.Query("occupancies") == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{"occupancies is required"})
		return
	}

	// Split the hotelIds query parameter into a slice of strings
	// and convert each string to an int64
	hotelIDChunks := strings.Split(c.Query("hotelIds"), ",")
	req.HotelIds = []int64{}
	for _, hotelID := range hotelIDChunks {
		id, err := strconv.ParseInt(hotelID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
			return
		}
		req.HotelIds = append(req.HotelIds, id)
	}

	// Unmarshal the occupancies query parameter into the Occupancies field
	if err := json.Unmarshal([]byte(c.Query("occupancies")), &req.Occupancies); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	// Call the app to get the hotel prices
	res, err := h.app.GetHotels(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
