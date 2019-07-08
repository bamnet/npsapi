// Package npsapi provides an API to access National Park Service (NPS) data.
package npsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://developer.nps.gov/api/v1/"
)

// Client connects to the NPS API backend.
type Client struct {
	httpClient *http.Client
	apiKey     string
}

// NewClient creates a new NPS API client.
//
// An API Key is required for use and can be obtained at:
// https://www.nps.gov/subjects/developer/get-started.htm
func NewClient(apiKey string) *Client {
	return &Client{&http.Client{}, apiKey}
}

func (c *Client) fetch(ctx context.Context, endpoint string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = u.Path + endpoint
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", c.apiKey)
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := parseError(body); err != nil {
		return nil, err
	}

	return body, nil
}

type apiError struct {
	Error struct {
		Code    string
		Message string
	}
}

// parseError attempts to parse an JSON-encoded error from the API response.
func parseError(body []byte) error {
	e := apiError{}
	err := json.Unmarshal(body, &e)
	if err != nil {
		return err
	}
	if (apiError{}) != e {
		return fmt.Errorf("%s - %s", e.Error.Code, e.Error.Message)
	}

	return nil
}
