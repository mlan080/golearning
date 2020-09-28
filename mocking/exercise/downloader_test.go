package exercise

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO create a fake implementation of the Cache interface
type FakeCache struct{ cache bool }

func (f *FakeCache) Get(key string) (err error, ok bool, value string) {
	return nil, f.cache, "value"
}

func (f *FakeCache) Store(key, value string) error {
	return nil
}
func TestDownloadCached(t *testing.T) {
	f := &FakeCache{true}
	downloader := New(f)
	want := "value"
	err, got := downloader.Download("value")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestDownloadNotCached(t *testing.T) {
	want := "value"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, want)
	}))
	defer ts.Close()

	f := &FakeCache{false}
	downloader := New(f)
	err, got := downloader.Download(ts.URL)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

// TODO implement Get and Store for MockCache

type MockCache struct {
	mock.Mock
}

func (c *MockCache) Get(key string) (error, bool, string) {
	args := c.Called(key)
	return args.Error(0), args.Get(1).(bool), args.Get(2).(string)
}

func (c *MockCache) Store(key, value string) error {
	args := c.Called(key, value)
	return args.Error(0)
}
func TestDownloadCacheUpdate(t *testing.T) {
	want := "value"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, want)
	}))
	defer ts.Close()

	m := &MockCache{}
	downloader := New(m)
	m.On("Get", ts.URL).Return(nil, bool(false), want)
	m.On("Store", ts.URL, "value").Return(nil)
	err, got := downloader.Download(ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
	m.AssertExpectations(t)
}
