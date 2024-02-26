package app

import (
	"encoding/json"
	"math/big"
	"nuitee-task/handlers"

	"go.uber.org/zap"
)

type App struct {
	hotelRates    HotelRates
	exchangeRates ExchangeRates
	log           *zap.SugaredLogger
}

func NewApp(hotelRates HotelRates, exchangeRates ExchangeRates, log *zap.SugaredLogger) *App {
	return &App{
		hotelRates:    hotelRates,
		exchangeRates: exchangeRates,
		log:           log,
	}
}

type HotelRates interface {
	GetHotelPrices(*handlers.GetHotelsRequest) ([]*handlers.HotelPrice, string, error)
}

type ExchangeRates interface {
	ConvertCurrency(amount *big.Float, from, to string) (*big.Float, error)
}

// GetHotels gets the hotel prices from the hotelRates service and converts them to the requested currency
func (a *App) GetHotels(req *handlers.GetHotelsRequest) (*handlers.GetHotelsResponse, error) {
	res := handlers.GetHotelsResponse{}

	hotelPrice, supplier, err := a.hotelRates.GetHotelPrices(req)
	if err != nil {
		a.log.Errorw("failed to get hotel prices", "error", err)
		return nil, err
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		a.log.Errorw("failed to marshal request", "error", err)
		return nil, err
	}

	res.Supplier.Request = string(reqBytes)
	res.Data, err = a.convertHotelPrices(hotelPrice, req.Currency)
	if err != nil {
		a.log.Errorw("failed to convert hotel prices", "error", err)
		return nil, err
	}
	res.Supplier.Response = supplier

	return &res, nil
}

func (a *App) Shutdown() {
	a.log.Info("shutting down app")
}

// convertHotelPrices converts the hotel prices to the requested currency
func (a *App) convertHotelPrices(prices []*handlers.HotelPrice, toCurrency string) ([]*handlers.HotelPrice, error) {
	if len(prices) == 0 {
		return prices, nil
	}

	for k, hPrice := range prices {
		if hPrice.Currency == toCurrency {
			continue
		}

		convertedAmount, err := a.exchangeRates.ConvertCurrency(hPrice.Price, hPrice.Currency, toCurrency)
		if err != nil {
			return nil, err
		}

		prices[k].Price = convertedAmount
		prices[k].Currency = toCurrency
	}

	return prices, nil
}
