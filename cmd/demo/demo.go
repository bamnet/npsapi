// demo application to exercise features of the API.
package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/bamnet/npsapi"
)

var (
	apiKey = flag.String("key", "", "API Key for the NPS API.")
	action = flag.String("action", "list", "Action to perform. One of 'list', 'get'.")
)

func init() {
	flag.Parse()
}

func main() {
	if *apiKey == "" {
		fmt.Println("An API key is required")
		return
	}

	client := npsapi.NewClient(*apiKey)
	ctx := context.Background()

	switch *action {
	case "list":
		parks, err := client.ListParks(ctx)
		if err != nil {
			fmt.Printf("ListParks error: %v", err)
			return
		}
		fmt.Printf("Fetched %d parks\n", len(parks))
		fmt.Printf("Park 0: %+v\n", parks[0])
	case "get":
		park, err := client.GetPark(ctx, "edis")
		if err != nil {
			fmt.Printf("GetPark error: %v", err)
			return
		}
		fmt.Printf("Park: %+v\n", park)
	default:
		fmt.Println("Unknown command.")
	}
}
