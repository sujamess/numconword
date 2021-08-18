package currency

type CoinCurrency string

var CoinCurrencies = map[Currency]CoinCurrency{
	USD: "cents",
	THB: "satangs",
}

func (c CoinCurrency) String() string {
	return string(c)
}
