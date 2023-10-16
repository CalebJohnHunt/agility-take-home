package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Same as getFromUrl, but takes a slice of URLs and
// returns a slice of T.
func getFromUrls[T any](urls []string) ([]T, error) {
	ts := make([]T, len(urls))
	// errChan acts as our wait group
	errChan := make(chan error)
	defer close(errChan)
	for i, url := range urls {
		// Get each url concurrently
		go func(i int, url string) {
			t, err := getFromUrl[T](url)
			if err != nil {
				errChan <- err
				return
			}

			ts[i] = t
			errChan <- nil
		}(i, url)
	}
	var retErr error = nil
	for i := 0; i < len(urls); i++ {
		if err := <-errChan; err != nil {
			retErr = errors.Join(err, retErr)
		}
	}
	return ts, nil
}

// Send an HTTP GET request to the given URL and
// unmarshall the response's JSON into type T.
// I believe T must not be a nilable type.
func getFromUrl[T any](url string) (T, error) {
	var t T
	resp, err := http.Get(url)
	if err != nil {
		return t, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return t, err
	}

	err = json.Unmarshal(b, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
