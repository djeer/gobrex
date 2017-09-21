package main

import (
"fmt"
"net/http"
"io/ioutil"
"encoding/json"
"time"
)

const api_url string = "https://bittrex.com/api/v1.1/public/getticker?market="

type ApiResponse struct {
	Success  bool
	Message  string
	Result   ResponseResultType
}

type ResponseResultType struct {
	//Bid    float64
	//Ask    float64
	Last   float64
}

func CheckCurrency(currency string, m map[string]float64) {
	resp, err := http.Get(api_url+currency)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	// parse json to struct
	var response ApiResponse
	json.Unmarshal(body, &response)
    // compare with old value
	if response.Result.Last != m[currency]{
		fmt.Print(time.Now().Format("15:04:05.99 "))
		fmt.Print(currency)
		fmt.Printf(" Last = %.8f", response.Result.Last)
		fmt.Printf(" was = %.8f\n", m[currency])
		// assign new value
		m[currency] = response.Result.Last
	}
}

func main() {
	// get config
	var args []string = []string{"BTC-ETH", "BTC-LTC", "BTC-XMR", "BTC-NXT", "BTC-DASH"}
	var delay = time.Second / 3
	// map to store values
	var m map[string]float64
	m = make(map[string]float64)
	// start an endless loop
	for i := 0;; i++ {
		if i == len(args) {
			i = 0
		}
		go CheckCurrency(args[i], m) // no time shifts here
		time.Sleep(delay)
	}
}
