package currency

type Currency string

const (
	USD Currency = "USD"
	THB Currency = "THB"
)

func (c *Currency) String() string {
	return string(*c)
}
