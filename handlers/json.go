package handlers

import (
	"encoding/json"
	"math"
	"math/big"
)

// GetHotelsRequest is the request for the GetHotels route
type GetHotelsRequest struct {
	CheckIn          string      `form:"checkin"`
	CheckOut         string      `form:"checkout"`
	Currency         string      `form:"currency"`
	GuestNationality string      `form:"guestNationality"`
	HotelIds         []int64     //custom validation and processing in the handler
	Occupancies      []Occupancy //custom validation and processing in the handler
}

type Occupancy struct {
	Adults   int64 `json:"adults"`
	Rooms    int64 `json:"rooms"`
	Children int64 `json:"children"`
}

// GetHotelsResponse is the response for the GetHotels route
type GetHotelsResponse struct {
	Data     []*HotelPrice `json:"data"`
	Supplier Supplier      `json:"supplier"`
}

type Supplier struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type HotelPrice struct {
	HotelID  string     `json:"hotelId"`
	Currency string     `json:"currency"`
	Price    *big.Float `json:"price"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// UnmarshalJSON unmarshals the HotelPrice from JSON
// It converts the price to a big.Float for test assertions
func (h *HotelPrice) UnmarshalJSON(data []byte) error {
	type Alias HotelPrice
	aux := &struct {
		Price float64 `json:"price"`
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	h.Price = big.NewFloat(aux.Price)

	return nil
}

// MarshalJSON marshals the HotelPrice to JSON
// It rounds the price to two decimal places as required by the spec
func (h *HotelPrice) MarshalJSON() ([]byte, error) {
	type Alias HotelPrice
	f, _ := h.Price.Float64()
	rounded := math.Round(f*100) / 100

	return json.Marshal(&struct {
		*Alias
		Price float64 `json:"price"`
	}{
		Price: rounded,
		Alias: (*Alias)(h),
	})
}
