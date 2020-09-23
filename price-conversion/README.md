# Price Conversion Exercise

This is a simplified version of something our backend does: convert prices for a trip from one
currency to another, so the user sees everything in their own currency.

Your task is to implement two different ways to do price conversion. One just applies the right
conversion rate. The other applies the conversion rate, and also adds a markup of 10%. We’ll use an
interface so we can use the two implementations interchangeably.


## Prices and conversion rates

Prices are always stored in the smaller unit, e.g. EUR 19.00 is stored as 1900. This way we can use
`int` for prices. (That doesn’t work for all currencies, but it’s fine for the ones we have here.)

The conversion rates are stored like this:

    var rates map[string]map[string]float64

Looks a bit confusing, but it’s not hard to use. For example,

    rates["EUR"]

has the conversion rates to convert from EUR to another currency, and

    rates["EUR"]["USD"]

is the rate to convert from EUR to USD. So you’d use something like

    amountInEUR * rates["EUR"]["USD"]

to convert an amount from EUR to USD.


## Running the program

The code here implements a command-line program with some hard-coded trips, prices and conversion
rates. You run it like this:

    go run cmd/main.go                         # get prices in EUR (the default)
    go run cmd/main.go -currency USD           # get prices in USD
    go run cmd/main.go -currency USD -markup   # get prices in USD, with 10% markup

The currencies you can use are EUR, USD, GBP, and CHF.

Look for the “Final price” values in the output -- you can see it doesn’t quite work yet. When
you’re done with the exercise, it should looks something like this:

    price-conversion$ go run cmd/main.go -currency USD -markup
    Trips:
     * Departure 2020-09-08T12:30:00Z Berlin
       Arrival   2020-09-08T13:50:00Z Hamburg
       Original price: EUR 40.00
       Final price:    USD 51.92
     * Departure 2020-09-10T08:10:00Z Zurich
       Arrival   2020-09-10T11:20:00Z Geneva
       Original price: CHF 85.00
       Final price:    USD 101.91
     * Departure 2020-09-15T09:00:00Z London
       Arrival   2020-09-15T13:10:00Z Edinburgh
       Original price: GBP 72.00
       Final price:    USD 103.75


## Files

Source files:

* `trips/trip.go` defines types for trips and prices. You don’t need to change anything here.
* `cmd/main.go` is the command-line program. The trips and prices are hard-coded here. Most of this
  is done, but you still need to complete the definition of the `Converter` interface and use the
  interface to convert prices.
* `simple/simple.go` implements simple currency conversion. It basically calculates `amount * rate`.
  This is up to you to implement.
* `markup/markup.go` implements currency conversion with markup. It basically calcuates
  `amount * rate * markup`. You can probably copy some of the “simple“ code, then add the markup
  logic.

All places in the code where you need to make changes are marked with “TODO.”


## Taking a step back

When you’re done with the exercise, take another look at the overall structure of the code. Which
package imports which other packages? Where is the `Converter` interface defined?

Notice that the packages that *implement* the interface don’t need to import the one that *defines*
the interface.
