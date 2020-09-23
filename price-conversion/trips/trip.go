package trips

import (
	"fmt"
	"io"
	"time"
)

// A Trip is a trip with times and cost.
type Trip struct {
	From, To                  string
	Start, End                time.Time
	OriginalPrice, FinalPrice *Price
}

// Price is a price with currency and amount. The amount is in cent/pence/etc.
type Price struct {
	Currency string
	Amount   int
}

// Print a readable multi-line description.
func (t *Trip) Print(out io.Writer) {
	fmt.Fprintf(out, " * Departure %s %s\n", t.Start.Format(time.RFC3339), t.From)
	fmt.Fprintf(out, "   Arrival   %s %s\n", t.End.Format(time.RFC3339), t.To)
	fmt.Fprintf(out, "   Original price: %s %.2f\n", t.OriginalPrice.Currency, float64(t.OriginalPrice.Amount)/100.0)
	fmt.Fprintf(out, "   Final price:    %s %.2f\n", t.FinalPrice.Currency, float64(t.FinalPrice.Amount)/100.0)
}
