package main

import (
	"fmt"

	"github.com/go-resty/resty"
)

func main() {
	resp, err := resty.R().Get("http://www.thomas-bayer.com/sqlrest/CUSTOMER/18/")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
}
