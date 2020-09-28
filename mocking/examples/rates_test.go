package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeRates1(t *testing.T) {
	// test ExchangeRates using the real service -- not a good unit test
	rates := NewExchangeRates("https://api.exchangeratesapi.io/latest")
	err, got := rates.Get("EUR", "USD")
	assert.NoError(t, err)
	assert.Equal(t, 1.1987, got)
}

func TestExchangeRates2(t *testing.T) {
	// test ExchangeRates with a fixed JSON response and httptest
	response := `
{
  "rates": {
    "USD": 1.2345
  },
  "base": "EUR",
  "date": "2020-09-01"
}
	`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	}))
	defer ts.Close()

	rates := NewExchangeRates(ts.URL)
	err, got := rates.Get("EUR", "USD")
	assert.NoError(t, err)
	assert.Equal(t, 1.2345, got)
}
