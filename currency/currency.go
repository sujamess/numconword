package currency

// Currency represents a currency for selecting an integer currency
// when the flag WithCurrency in options.ConvertOptions is true
type Currency string

const (
	// USD represents a U.S. Dollar currency
	USD Currency = "USD"
	// THB represents a Thai Baht currency
	THB Currency = "THB"
)

// Point uses when the flag WithCurrency in options.ConvertOptions is false
const Point = "point"

// String gets a string value
func (c *Currency) String() string {
	return string(*c)
}
