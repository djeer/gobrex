package lib

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"errors"
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

func CheckCurrency(currency *string, m map[string]float64) (err error) {
	resp, err := http.Get(api_url+*currency)
	if err != nil {
		err := errors.New("could not fetch data")
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := errors.New("could not fetch data")
		fmt.Println(err)
		return err
	}
	// parse json to struct
	var response ApiResponse
	json.Unmarshal(body, &response)
	if response.Success != true {
		err := errors.New("API error: " + response.Message + " market: " + *currency)
		fmt.Println(err)
		return err
	}
	// compare with old value
	if response.Result.Last != m[*currency]{
		fmt.Print(time.Now().Format("15:04:05.99 "))
		fmt.Print(*currency)
		fmt.Printf(" Last = %.8f ", response.Result.Last)
		// print some arrows
		switch {
		case m[*currency] == 0: // first iteration
			fmt.Println()
		case response.Result.Last < m[*currency]:
			fmt.Println("↓")
		default:
			fmt.Println("↑")
		}
		// assign new value
		m[*currency] = response.Result.Last
	}
	return nil
}

