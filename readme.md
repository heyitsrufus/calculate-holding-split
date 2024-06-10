# Calculate Crypto Holding CLI

Calculate Crypto Holding is a command-line interface (CLI) tool that calculates the 70/30 split for two given cryptocurrencies based on the specified USD amount.

## Features

- Fetches real-time cryptocurrency exchange rates from the Coinbase API.
- Calculates the 70/30 split for the specified USD amount and cryptocurrencies.
- Supports custom split percentages and multiple cryptocurrencies.
- Easily upgradable with additional user input variables to do more calculation(s) or add more crypto currencies.
- Provides detailed error messages and logging.

## Requirements

- `encoding/json`: For decoding JSON responses from the API.
- `fmt`: For formatted I/O operations.
- `net/http`: For making HTTP requests to the Coinbase API.
- `os`: For handling command-line arguments other OS actions.
- `strconv`: For converting string representations of numbers to float64.

## Sample Usage

Open the project and run these commands from the project's root directory

```
go build -o calculate-crypto
./calculate-crypto 100 BTC ETH
```
