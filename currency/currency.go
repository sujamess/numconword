package currency

type Currency string

const (
	USD Currency = "USD"
	THB Currency = "THB"
)

const Point = "point"

// String gets a string value
func (c *Currency) String() string {
	return string(*c)
}
