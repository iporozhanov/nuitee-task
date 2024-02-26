package exchangerates

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"
)

type CoinbaseResponse struct {
	Data CoinbaseCurrencies `json:"data"`
}

type CoinbaseCurrencies struct {
	Currency string                `json:"currency"`
	Rates    map[string]*big.Float `json:"rates"`
}

type CoinbaseClient struct {
	Rates          map[string]map[string]*big.Float
	RatesClearTime time.Duration
	BaseURL        string
}

func NewCoinbaseClient(baseURL string, clearTime time.Duration) *CoinbaseClient {
	return &CoinbaseClient{
		Rates:          make(map[string]map[string]*big.Float),
		RatesClearTime: clearTime,
		BaseURL:        baseURL,
	}
}

// GetLatestRates gets the latest exchange rates from the Coinbase API.
func (c *CoinbaseClient) GetLatestRates(base string) (map[string]*big.Float, error) {
	resp, err := http.Get(fmt.Sprintf("%s/v2/exchange-rates?currency=%s", c.BaseURL, base))
	if err != nil {
		return nil, fmt.Errorf("failed to get latest rates: %w", err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var data CoinbaseResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return data.Data.Rates, nil
}

// ConvertCurrency converts the amount from one currency to another.
func (c *CoinbaseClient) ConvertCurrency(amount *big.Float, from, to string) (*big.Float, error) {
	// Check if the rate is already cached.
	if _, ok := c.Rates[from]; !ok {
		rates, err := c.GetLatestRates(from)
		if err != nil {
			return nil, err
		}
		// Cache the rates
		c.Rates[from] = rates
	}

	toRate := c.Rates[from][to]

	return amount.Mul(amount, toRate), nil
}

// ClearRates clears the exchange rates cache at regular intervals.
func (c *CoinbaseClient) ClearRates() {
	ticker := time.NewTicker(c.RatesClearTime)
	for range ticker.C {
		c.Rates = make(map[string]map[string]*big.Float)
	}
}
