// Package markup implement currency conversion with a markup.
package markup

// Converter converts from one currency to another.
type MarkupConverter struct {
	rate map[string]map[string]float64
	markup float64
}

// TODO Converter should implement the interface in main.go
func (c MarkupConverter) Convert(fromCurrency, toCurrency string, amount float64)float64{
	if fromCurrency == toCurrency {
		return amount 
	} 
	return amount * c.rate[fromCurrency][toCurrency] * c.markup
}

// New creates a new Converter with the given conversion rates.
func New(rates map[string]map[string]float64, markup float64) *MarkupConverter {
	return &MarkupConverter{rates, markup}
}
