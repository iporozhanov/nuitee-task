package hotelrates

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"nuitee-task/handlers"
	"time"
)

type GetRatesRequest struct {
	Stay        Stay             `json:"stay"`
	Occupancies []Occupancy      `json:"occupancies"`
	Hotels      HotelRequestList `json:"hotels"`
}

type Stay struct {
	CheckIn  string `json:"checkIn"`
	CheckOut string `json:"checkOut"`
}

type Occupancy struct {
	Adults   int64 `json:"adults"`
	Rooms    int64 `json:"rooms"`
	Children int64 `json:"children"`
}

type HotelRequestList struct {
	Hotel []int64 `json:"hotel"`
}

type GetRatesResponse struct {
	AuditData any           `json:"auditData"`
	Hotels    HotelResponse `json:"hotels"`
}

type HotelResponse struct {
	Hotels   []Hotel `json:"hotels"`
	CheckIn  string  `json:"checkIn"`
	CheckOut string  `json:"checkOut"`
	Total    int64   `json:"total"`
}

type Hotel struct {
	Code            int64      `json:"code"`
	Name            string     `json:"name"`
	ExclusiveDeal   int64      `json:"exclusiveDeal"`
	CategoryCode    string     `json:"categoryCode"`
	CategoryName    string     `json:"categoryName"`
	DestinationCode string     `json:"destinationCode"`
	DestinatioName  string     `json:"destinationName"`
	ZoneCode        int64      `json:"zoneCode"`
	ZoneName        string     `json:"zoneName"`
	Latitude        string     `json:"latitude"`
	Longitude       string     `json:"longitude"`
	Rooms           []any      `json:"rooms"` //we want this in the response string, but don't need to use any of the data
	MinRate         *big.Float `json:"minRate"`
	MaxRate         *big.Float `json:"maxRate"`
	Currency        string     `json:"currency"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type HotelbedsClient struct {
	APIKey string
	Secret string
	APIUrl string
}

func NewHotelbedsClient(apiKey, secret, apiUrl string) *HotelbedsClient {
	return &HotelbedsClient{
		APIKey: apiKey,
		Secret: secret,
		APIUrl: apiUrl,
	}
}

// xSignature generates the x-signature header value
func (c *HotelbedsClient) xSignature() string {
	now := time.Now().Unix()

	combined := fmt.Sprintf("%s%s%d", c.APIKey, c.Secret, now)

	h := sha256.New()
	h.Write([]byte(combined))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// fetchRates sends a request to the hotelbeds API and returns the response
func (c *HotelbedsClient) fetchRates(req *GetRatesRequest) (*GetRatesResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %w", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.APIUrl, "hotel-api/1.0/hotels"), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Api-key", c.APIKey)
	request.Header.Set("X-Signature", c.xSignature())

	// Send the request
	client := &http.Client{}
	r, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error getting hotel rates: %w", err)
	}

	// Read the response
	responseBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if r.StatusCode != http.StatusOK {
		errResp := &ErrorResponse{}
		err = json.Unmarshal(responseBody, errResp)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling error response: %w", err)
		}

		return nil, fmt.Errorf("error getting hotel rates: %s", errResp.Error)
	}

	resp := &GetRatesResponse{}

	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return resp, nil
}

// GetHotelPrices returns the prices for the given hotels
func (c *HotelbedsClient) GetHotelPrices(req *handlers.GetHotelsRequest) ([]*handlers.HotelPrice, string, error) {
	// Convert the request to the hotelbeds request
	occupancies := make([]Occupancy, len(req.Occupancies))
	for i, o := range req.Occupancies {
		occupancies[i] = Occupancy{
			Adults:   o.Adults,
			Rooms:    o.Rooms,
			Children: o.Children,
		}
	}

	// Fetch the rates
	resp, err := c.fetchRates(&GetRatesRequest{
		Stay: Stay{
			CheckIn:  req.CheckIn,
			CheckOut: req.CheckOut,
		},
		Occupancies: occupancies,
		Hotels: HotelRequestList{
			Hotel: req.HotelIds,
		},
	})
	if err != nil {
		return nil, "", err
	}

	// Convert the response to the expected format
	hotelPrices := make([]*handlers.HotelPrice, len(resp.Hotels.Hotels))
	for i, h := range resp.Hotels.Hotels {

		hotelPrices[i] = &handlers.HotelPrice{
			Price:    h.MinRate,
			HotelID:  fmt.Sprintf("%d", h.Code),
			Currency: h.Currency,
		}
	}

	// Convert the hotelbeds response to a string
	bResponse, err := json.Marshal(resp)
	if err != nil {
		return nil, "", err
	}

	return hotelPrices, string(bResponse), nil
}
