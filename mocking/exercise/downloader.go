package exercise

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Cache implements a way to cache strings.
type Cache interface {
	// Get retrieves a value stored under a key. If the key can’t be found, it returns
	// ok==false.
	Get(key string) (err error, ok bool, value string)

	// Store stores a value under a key.
	Store(key, value string) error
}

// Downloader downloads files over HTTP using some kind of cache.
type Downloader struct {
	cache Cache
}

// New creates a new Downloader.
func New(cache Cache) *Downloader {
	return &Downloader{cache}
}

// Download retrieves a URL and updates the cache.
func (d *Downloader) Download(url string) (error, string) {
	// look up URL in cache
	err, ok, value := d.cache.Get(url)
	if err != nil {
		return fmt.Errorf("error in cache lookup: %v", err), ""
	}
	if ok {
		return nil, value
	}

	// fetch URL
	resp, err := http.Get("url")
	if err != nil {
		return fmt.Errorf("error downloading: %v", err), ""
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error downloading: %v", err), ""
	}
	value = string(data)

	// update cache
	err = d.cache.Store(url, value)
	if err != nil {
		return fmt.Errorf("error in cache update: %v", err), ""
	}

	return nil, value
}
