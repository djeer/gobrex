package main

import (
	"testing"
	"./lib"
)

func TestCheckCurrency(t *testing.T) {

	// map init
	var m map[string]float64
	m = make(map[string]float64)

	var invalid_currency string = "VVV-XXX"
	lib.CheckCurrency(&invalid_currency, m)
	// Output: API error: INVALID_MARKET market: VVV-XXX
	if m[invalid_currency] != 0.0 {
		t.Fatal("parsed non-existing currency")
	}

	var valid_currency string = "BTC-ETH"
	lib.CheckCurrency(&valid_currency, m)
	if m[valid_currency] == 0.0 {
		t.Fatal("cannot parse existing currency")
	}
}
