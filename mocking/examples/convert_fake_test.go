package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FakeRates is a "fake" implementation of the Rates interface.
type FakeRates struct{}

func (f *FakeRates) Get(from, to string) (error, float64) {
	return nil, 2
}

func TestConvertWithFake(t *testing.T) {
	// unit test for Convert using a fake
	input := 1.23
	want := 2 * input
	err, got := Convert(new(FakeRates), "EUR", "USD", input)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
