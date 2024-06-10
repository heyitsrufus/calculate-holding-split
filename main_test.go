package main

import (
	"testing"
)

/*
This tests the calculateHoldingSplit function.
This function takes in a USD amount and two crypto symbols with their exchange rates,
calculates the 70/30 split, and returns the respective amounts of each cryptocurrency.
*/
func TestCalculateHoldingSplit(t *testing.T) {

	// Setting sample input values
	usdAmount := 100.0
	crypto1 := "BTC"
	crypto2 := "ETH"
	rate1 := "0.00005"
	rate2 := "0.002"

	amount1, amount2, err := calculateHoldingSplit(usdAmount, crypto1, crypto2, rate1, rate2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Expected values based on the input we set
	expectedAmount1 := (usdAmount * 0.70) * 0.00005
	expectedAmount2 := (usdAmount * 0.30) * 0.002

	if amount1 != expectedAmount1 {
		t.Errorf("expected %f, got %f", expectedAmount1, amount1)
	}

	if amount2 != expectedAmount2 {
		t.Errorf("expected %f, got %f", expectedAmount2, amount2)
	}
}
