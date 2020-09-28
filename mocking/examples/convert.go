package example

// Rates is an interface for retrieving currency conversion rates.
type Rates interface {
	Get(to, from string) (error, float64)
}

// Convert uses a Rates object to convert from one currency to another.
func Convert(r Rates, to, from string, amount float64) (error, float64) {
	err, rate := r.Get(to, from)
	if err != nil {
		return err, 0
	}
	return nil, rate * amount
}
