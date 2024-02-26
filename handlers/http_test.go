package handlers_test

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"nuitee-task/handlers"
	"nuitee-task/handlers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHotels(t *testing.T) {
	app := mocks.NewAPP(t)
	h := handlers.NewHTTP(app)
	h.InitRoutes()
	expectedResponse := &handlers.GetHotelsResponse{
		Data: []*handlers.HotelPrice{
			{
				HotelID:  "1",
				Currency: "USD",
				Price:    big.NewFloat(100),
			},
			{
				HotelID:  "2",
				Currency: "USD",
				Price:    big.NewFloat(200),
			},
		},
		Supplier: handlers.Supplier{
			Request:  "request",
			Response: "response",
		},
	}

	app.EXPECT().GetHotels(&handlers.GetHotelsRequest{
		CheckIn:  "2024-03-15",
		CheckOut: "2024-03-16",
		Currency: "USD",
		HotelIds: []int64{1, 2, 3},
		Occupancies: []handlers.Occupancy{
			{Rooms: 2, Adults: 2},
			{Rooms: 1, Adults: 1},
		},
	}).Return(expectedResponse, nil)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/hotels?hotelIds=1,2,3&currency=USD&checkin=2024-03-15&checkout=2024-03-16&occupancies=[{\"rooms\":2,\"adults\":2},{\"rooms\":1,\"adults\":1}]", nil)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Perform the request
	h.Router.ServeHTTP(res, req)

	// Assert that the response status code is 200
	assert.Equal(t, http.StatusOK, res.Code)

	// Assert that the response body is not empty
	assert.NotEmpty(t, res.Body.String())

	// Parse the response body into a JSON object
	var response handlers.GetHotelsResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse.Data[0].Price, response.Data[0].Price)
	assert.Equal(t, expectedResponse.Data[1].Price, response.Data[1].Price)
	assert.Equal(t, expectedResponse.Supplier.Request, response.Supplier.Request)
	assert.Equal(t, expectedResponse.Supplier.Response, response.Supplier.Response)

	app.AssertExpectations(t)
}

func TestGetHotels_BadRequest(t *testing.T) {
	app := mocks.NewAPP(t)
	h := handlers.NewHTTP(app)
	h.InitRoutes()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/hotels", nil)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Perform the request
	h.Router.ServeHTTP(res, req)

	// Assert that the response status code is 400
	assert.Equal(t, http.StatusBadRequest, res.Code)

	// Assert that the response body is not empty
	assert.NotEmpty(t, res.Body.String())

	// Parse the response body into a JSON object
	var response handlers.ErrorResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "hotelIds is required", response.Error)

	app.AssertExpectations(t)
}

func TestGetHotels_BadRequest2(t *testing.T) {
	app := mocks.NewAPP(t)
	h := handlers.NewHTTP(app)
	h.InitRoutes()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/hotels?hotelIds=1,2,3", nil)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Perform the request
	h.Router.ServeHTTP(res, req)

	// Assert that the response status code is 400
	assert.Equal(t, http.StatusBadRequest, res.Code)

	// Assert that the response body is not empty
	assert.NotEmpty(t, res.Body.String())

	// Parse the response body into a JSON object
	var response handlers.ErrorResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "occupancies is required", response.Error)

	app.AssertExpectations(t)
}

func TestGetHotels_AppError(t *testing.T) {
	app := mocks.NewAPP(t)
	h := handlers.NewHTTP(app)
	h.InitRoutes()

	expectedError := assert.AnError

	app.EXPECT().GetHotels(&handlers.GetHotelsRequest{
		CheckIn:  "2024-03-15",
		CheckOut: "2024-03-16",
		Currency: "USD",
		HotelIds: []int64{1, 2, 3},
		Occupancies: []handlers.Occupancy{
			{Rooms: 2, Adults: 2},
			{Rooms: 1, Adults: 1},
		},
	}).Return(nil, expectedError)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/hotels?hotelIds=1,2,3&currency=USD&checkin=2024-03-15&checkout=2024-03-16&occupancies=[{\"rooms\":2,\"adults\":2},{\"rooms\":1,\"adults\":1}]", nil)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Perform the request
	h.Router.ServeHTTP(res, req)

	// Assert that the response status code is 500
	assert.Equal(t, http.StatusInternalServerError, res.Code)

	// Assert that the response body is not empty
	assert.NotEmpty(t, res.Body.String())

	// Parse the response body into a JSON object
	var response handlers.ErrorResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedError.Error(), response.Error)

	app.AssertExpectations(t)
}
