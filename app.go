package main

import (
	"time"
	"./lib"
)

func main() {
	// get config
	var args []string = []string{"BTC-ETH", "BTC-LTC", "BTC-XMR", "BTC-NXT", "BTC-DASH"}
	var delay = time.Second / 3
	// map to store values
	var m map[string]float64
	m = make(map[string]float64)
	// start an endless loop
	for {
	    for i := 0;; i++ {
		    if i == len(args) {
			    break
		    }
		go lib.CheckCurrency(&args[i], m)
		}
		time.Sleep(delay)  // no time shifts here
	}
}
