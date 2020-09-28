package example

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ExchangeRates downloads currency exchange rates from exchangeratesapi.io.
type ExchangeRates struct {
	url string
}

// NewExchangeRates creates a new ExchangeRates object.
func NewExchangeRates(url string) *ExchangeRates {
	return &ExchangeRates{url}
}

// Get fetches the rate to convert from "from" to "to", for example, from EUR to GBP.
func (r *ExchangeRates) Get(from, to string) (error, float64) {
	// fetch JSON
	resp, err := http.Get(fmt.Sprintf("%s?base=%s&symbols=%s", r.url, from, to))
	if err != nil {
		return fmt.Errorf("error fetching conversion rate from exchangeratesapi: %w", err), 0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response from exchangeratesapi: %w", err), 0
	}

	// parse JSON and get conversion rate
	var data struct {
		Rates map[string]float64
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("error parsing response from exchangeratesapi: %w", err), 0
	}
	return nil, data.Rates[to]
}
