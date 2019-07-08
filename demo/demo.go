package main

import (
	"fmt"
	"log"

	npsclient "github.com/bamnet/npsapi/generated/client"
	oapiclient "github.com/go-openapi/runtime/client"
)

func main() {
	fmt.Print("hi")

	c := npsclient.Default

	auth := oapiclient.APIKeyAuth("X-Api-Key", "header", "api-key-here")

	resp, err := c.Parks.GetPark(nil, auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response: %v", resp.Payload)
}
