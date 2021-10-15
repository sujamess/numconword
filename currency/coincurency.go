package currency

type CoinCurrency string

var CoinCurrencies = map[Currency]CoinCurrency{
	USD: "cents",
	THB: "satangs",
}

// String gets a string value
func (c CoinCurrency) String() string {
	return string(c)
}
