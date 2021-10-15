package currency

// CoinCurrency represents a currency for selecting a decimal currency
// when the flag WithCurrency in options.ConvertOptions is true.
type CoinCurrency string

// CoinCurrencies is a map between currency and coin currency
var CoinCurrencies = map[Currency]CoinCurrency{
	USD: "cents",
	THB: "satangs",
}

// String gets a string value
func (c CoinCurrency) String() string {
	return string(c)
}
