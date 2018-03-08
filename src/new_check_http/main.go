package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)

func main() {
	url := flag.String("nombre", "", "Name to salute.")
	status := flag.Int("status", 0, "Expected HTTP status.")
	flag.Parse()

	fmt.Printf("Url to check: %s (type %T)\n", *url, *url)
	fmt.Printf("Status to check: %i (type %T)\n", *status, *status)
	result, err := checkHttpCode(*url, *status)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Check result is ", result)
}

func checkHttpCode(url string, expected_status int) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	fmt.Println(resp)
	return resp.StatusCode == expected_status, nil
}
