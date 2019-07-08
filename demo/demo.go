package main

import (
	"context"
	"fmt"

	"github.com/bamnet/npsapi"
)

func main() {
	client := npsapi.NewClient("api-key-here")
	parks, err := client.ListParks(context.Background())
	if err != nil {
		fmt.Printf("ListParks error: %v", err)
	}
	fmt.Printf("Parks: %v+", parks)
}
