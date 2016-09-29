package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

// JSONTime IS THE TIME
type JSONTime struct { // this is a time struct
	Time  string `json:"time"`
	Epoch string `json:"millisecond_since_epoch"`
	Date  string `json:"date"`
}

func main() {
	resp, err := resty.R().Get("http://time.jsontest.com/")

	if err != nil {
		fmt.Printf("\nError %v", err)
	}

	var timeVar JSONTime
	fmt.Print(resp.String())
	err2 := json.Unmarshal(resp.Body(), timeVar)

	if err2 != nil {
		fmt.Printf("Unmarshalling error %v\n", err)
	}
	fmt.Printf("Response Body: %+v", timeVar)
}
