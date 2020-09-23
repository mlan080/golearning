package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fromAtoB/engineering/training/backend/1/price-conversion/markup"
	"github.com/fromAtoB/engineering/training/backend/1/price-conversion/simple"
	"github.com/fromAtoB/engineering/training/backend/1/price-conversion/trips"
)

// Converter is the interface for an object that can do currency conversion.
type Converter interface {
	Convert(fromCurrency string, toCurrency string, amount float64) float64
}

func main() {
	// command-line flags
	currency := flag.String("currency", "EUR", "three-letter currency code")
	useMarkup := flag.Bool("markup", false, "apply 10% markup")
	flag.Parse()

	// trips and rates are hard-coded
	ts := sampleTrips()
	rates := sampleRates()

	// create converter
	var converter Converter
	if *useMarkup {
		converter = markup.New(rates, 1.10)
	} else {
		converter = simple.New(rates)
	}

	// convert prices
	for _, t := range ts {
		originalPrice := t.OriginalPrice

		// TODO use "converter" to convert "originalPrice" to "currency"
		convertedAmount := converter.Convert(originalPrice.Currency, *currency, float64(originalPrice.Amount))
		_ = originalPrice
		 _ = converter
		 _ = currency

		t.FinalPrice = &trips.Price{
			Currency: *currency,
			Amount:   int(convertedAmount),
		}
	}

	// print trips
	fmt.Println("Trips:")
	for _, t := range ts {
		t.Print(os.Stdout)
	}
}

func sampleTrips() []*trips.Trip {
	return []*trips.Trip{
		{
			From:  "Berlin",
			To:    "Hamburg",
			Start: time.Date(2020, 9, 8, 12, 30, 0, 0, time.UTC),
			End:   time.Date(2020, 9, 8, 13, 50, 0, 0, time.UTC),
			OriginalPrice: &trips.Price{
				Currency: "EUR",
				Amount:   4000,
			},
		},
		{
			From:  "Zurich",
			To:    "Geneva",
			Start: time.Date(2020, 9, 10, 8, 10, 0, 0, time.UTC),
			End:   time.Date(2020, 9, 10, 11, 20, 0, 0, time.UTC),
			OriginalPrice: &trips.Price{
				Currency: "CHF",
				Amount:   8500,
			},
		},
		{
			From:  "London",
			To:    "Edinburgh",
			Start: time.Date(2020, 9, 15, 9, 0, 0, 0, time.UTC),
			End:   time.Date(2020, 9, 15, 13, 10, 0, 0, time.UTC),
			OriginalPrice: &trips.Price{
				Currency: "GBP",
				Amount:   7200,
			},
		},
	}
}

func sampleRates() map[string]map[string]float64 {
	return map[string]map[string]float64{
		"EUR": {
			"USD": 1.18,
			"GBP": 0.90,
			"CHF": 1.08,
		},
		"USD": {
			"EUR": 0.85,
			"GBP": 0.77,
			"CHF": 0.92,
		},
		"GBP": {
			"EUR": 1.11,
			"USD": 1.31,
			"CHF": 1.20,
		},
		"CHF": {
			"EUR": 0.93,
			"USD": 1.09,
			"GBP": 0.84,
		},
	}
}
