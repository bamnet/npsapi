package npsapi

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
)

// ErrParkNotFound is a type of error returned when no parks were found.
var ErrParkNotFound = errors.New("park not found")

// Park basics include location, contact, and operating hours for each national park.
type Park struct {
	// UUID-style identifier for the park.
	ID string
	// Park name.
	Name string
	// Longer park name.
	FullName string
	// 4-character code to identify the park.
	Code string
	// States the park is in.
	States []string
}

// ListParks finds all the parks.
func (c *Client) ListParks(ctx context.Context) ([]Park, error) {
	params := map[string]string{
		"limit": "99999", // Fetch more parks than we have.
	}
	return c.fetchParks(ctx, params)
}

// GetPark finds a specific park with a 4-character code.
func (c *Client) GetPark(ctx context.Context, code string) (*Park, error) {
	params := map[string]string{
		"parkCode": code,
		"limit":    "1",
	}

	park, err := c.fetchParks(ctx, params)
	if err != nil {
		return nil, err
	}
	if len(park) >= 1 {
		return &park[0], nil
	}
	return nil, ErrParkNotFound
}

func (c *Client) fetchParks(ctx context.Context, params map[string]string) ([]Park, error) {
	resp, err := c.fetch(ctx, "parks", params)
	if err != nil {
		return nil, err
	}

	data := struct {
		Data []struct {
			ID       string
			Name     string
			FullName string
			ParkCode string
			States   string
		}
	}{}

	if err := json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	parks := []Park{}
	for _, p := range data.Data {
		park := Park{
			ID:       p.ID,
			Name:     p.Name,
			FullName: p.FullName,
			Code:     p.ParkCode,
			States:   strings.Split(p.States, ","),
		}
		parks = append(parks, park)
	}

	return parks, nil
}
