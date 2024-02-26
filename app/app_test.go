package app_test

import (
	"errors"
	"math/big"
	"nuitee-task/app"
	"nuitee-task/app/mocks"
	"nuitee-task/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetHotels_Success(t *testing.T) {
	hotelRates := mocks.NewHotelRates(t)
	exchangeRates := mocks.NewExchangeRates(t)
	logger := zap.NewNop().Sugar()

	app := app.NewApp(hotelRates, exchangeRates, logger)

	req := &handlers.GetHotelsRequest{
		CheckIn:  "2024-03-15",
		CheckOut: "2024-03-16",
		Currency: "USD",
		HotelIds: []int64{1, 2, 3},
		Occupancies: []handlers.Occupancy{
			{Rooms: 2, Adults: 2},
			{Rooms: 1, Adults: 1},
		},
	}

	expectedHotelPrices := []*handlers.HotelPrice{
		{
			HotelID:  "1",
			Currency: "EUR",
			Price:    big.NewFloat(100),
		},
		{
			HotelID:  "2",
			Currency: "EUR",
			Price:    big.NewFloat(200),
		},
	}

	expectedSupplier := "supplier"

	hotelRates.On("GetHotelPrices", req).Return(expectedHotelPrices, expectedSupplier, nil)
	exchangeRates.On("ConvertCurrency", big.NewFloat(100), "EUR", "USD").Return(big.NewFloat(100), nil)
	exchangeRates.On("ConvertCurrency", big.NewFloat(200), "EUR", "USD").Return(big.NewFloat(200), nil)

	response, err := app.GetHotels(req)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expectedHotelPrices, response.Data)
	assert.Equal(t, expectedSupplier, response.Supplier.Response)

	hotelRates.AssertExpectations(t)
	exchangeRates.AssertExpectations(t)
}

func TestGetHotels_ErrorGettingHotelPrices(t *testing.T) {
	hotelRates := mocks.NewHotelRates(t)
	exchangeRates := mocks.NewExchangeRates(t)
	logger := zap.NewNop().Sugar()

	app := app.NewApp(hotelRates, exchangeRates, logger)

	req := &handlers.GetHotelsRequest{
		CheckIn:  "2024-03-15",
		CheckOut: "2024-03-16",
		Currency: "USD",
		HotelIds: []int64{1, 2, 3},
		Occupancies: []handlers.Occupancy{
			{Rooms: 2, Adults: 2},
			{Rooms: 1, Adults: 1},
		},
	}

	expectedError := errors.New("failed to get hotel prices")

	hotelRates.On("GetHotelPrices", req).Return(nil, "", expectedError)

	response, err := app.GetHotels(req)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, expectedError, err)

	hotelRates.AssertExpectations(t)
	exchangeRates.AssertExpectations(t)
}

func TestGetHotels_ErrorConvertingCurrency(t *testing.T) {
	hotelRates := mocks.NewHotelRates(t)
	exchangeRates := mocks.NewExchangeRates(t)
	logger := zap.NewNop().Sugar()

	app := app.NewApp(hotelRates, exchangeRates, logger)

	req := &handlers.GetHotelsRequest{
		CheckIn:  "2024-03-15",
		CheckOut: "2024-03-16",
		Currency: "USD",
		HotelIds: []int64{1, 2, 3},
		Occupancies: []handlers.Occupancy{
			{Rooms: 2, Adults: 2},
			{Rooms: 1, Adults: 1},
		},
	}

	hotelPrices := []*handlers.HotelPrice{
		{
			HotelID:  "1",
			Currency: "EUR",
			Price:    big.NewFloat(100),
		},
		{
			HotelID:  "2",
			Currency: "EUR",
			Price:    big.NewFloat(200),
		},
	}

	expectedError := errors.New("failed to convert currency")

	hotelRates.On("GetHotelPrices", req).Return(hotelPrices, "", nil)
	exchangeRates.On("ConvertCurrency", big.NewFloat(100), "EUR", "USD").Return(big.NewFloat(100), nil)
	exchangeRates.On("ConvertCurrency", big.NewFloat(200), "EUR", "USD").Return(nil, expectedError)

	response, err := app.GetHotels(req)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, expectedError, err)

	hotelRates.AssertExpectations(t)
	exchangeRates.AssertExpectations(t)
}
