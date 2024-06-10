package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ExchangeRate struct {
	Rates map[string]string `json:"rates"`
}

// Function to fetch exchange rates from the Coinbase API.
func fetchCurrentCryptoExchangeRate() (map[string]string, error) {
	url := "https://api.coinbase.com/v2/exchange-rates?currency=USD"
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to retrive exchange rate: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the response returned a non-200 response code: %w", err)
	}

	var result struct {
		Data ExchangeRate `json:"data"`
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response from the Coinbase API response %w", err)
	}

	return result.Data.Rates, nil
}

func calculateHoldingSplit(usdAmount float64, crypto1 string, crypto2 string, rate1 string, rate2 string) (float64, float64, error) {

	//Convert the rates to float values
	rate1Float, err1 := strconv.ParseFloat(rate1, 64)
	if err1 != nil {
		return 0, 0, fmt.Errorf("error parsing rate for %s: %w", crypto1, err1)
	}

	rate2Float, err2 := strconv.ParseFloat(rate2, 64)
	if err2 != nil {
		return 0, 0, fmt.Errorf("error parsing rate for %s: %w", crypto2, err2)
	}

	// Calculate the 70 and 30 percent of the retrieved rate
	amount1 := (usdAmount * 0.70) * rate1Float
	amount2 := (usdAmount * 0.30) * rate2Float

	return amount1, amount2, nil
}

func main() {

	// Check if there is the correct user input
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./calculate-crypto <USD_AMOUNT> <CRYPTO #1> <CRYPTO #2>")
		os.Exit(1)
	}

	// Parse the command line arguments
	// Ensure the USD amount is numerical
	usdAmount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Invalid USD amount: %v\n", err)
		os.Exit(1)
	}

	// Get the two crypto from the command line arguments
	crypto1 := os.Args[2]
	crypto2 := os.Args[3]

	// Get the live crypto exchange rates
	fmt.Println("Fetching exchange rates...")
	rates, err := fetchCurrentCryptoExchangeRate()
	if err != nil {
		fmt.Printf("Error fetching exchange rates: %v\n", err)
		os.Exit(1)
	}

	rate1, ok1 := rates[crypto1]
	rate2, ok2 := rates[crypto2]

	// Check whether the user inputted cryptocurrency symbols in the Coinbase API response
	if !ok1 || !ok2 {
		fmt.Println("Can't find the user input cryptocurrency symbols in the Coinbase API response")
		os.Exit(1)
	}

	// Calculate the holding split and print out the result
	fmt.Println("Calculating the Holding Split...")
	amount1, amount2, err := calculateHoldingSplit(usdAmount, crypto1, crypto2, rate1, rate2)
	if err != nil {
		fmt.Printf("Error calculating the holding split: %v\n\n", err)
	}

	// Print out the results in the expected format
	fmt.Println("Here is the current 70/30 holding split for the specified crypto symbols:")
	fmt.Println()
	fmt.Printf("$%.2f => %.4f %s\n", usdAmount*0.70, amount1, crypto1)
	fmt.Printf("$%.2f => %.4f %s\n", usdAmount*0.30, amount2, crypto2)

}
