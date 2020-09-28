package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRates is a mock implementation of the Rates interface.
type MockRates struct {
	mock.Mock
}

func (r *MockRates) Get(from, to string) (error, float64) {
	args := r.Called(from, to)
	return args.Error(0), args.Get(1).(float64)
}

func TestConvertWithMock(t *testing.T) {
	// unit test for Convert using a mock object
	input := 1.23
	want := 2 * input
	rates := new(MockRates)
	rates.On("Get", "EUR", "USD").Return(nil, float64(2))
	err, got := Convert(rates, "EUR", "USD", input)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
	rates.AssertExpectations(t)
}
