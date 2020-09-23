// Package simple implements straightforward currency conversion.
package simple

// Converter converts from one currency to another.
type SimpleConverter struct {
	rates map[string]map[string]float64
}

// TODO Converter should implement the interface in main.go

// New creates a new Converter with the given conversion rates
//is this New to be to be used in main as .New? and has nothing to do with below
func New(conversionrates map[string]map[string]float64) *SimpleConverter {
s := SimpleConverter{}
s.rates = conversionrates //storing conversion rate in s 
return &s 
//return &SimpleConverter{conversionrates}
}




//implementing the method convert then use the field rate
func (c *SimpleConverter) Convert(fromCurrency, toCurrency string, amount float64) float64 {
if fromCurrency == toCurrency {
	return amount 
}
return amount * c.rates[fromCurrency][toCurrency]
}


