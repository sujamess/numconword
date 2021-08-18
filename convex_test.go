package convex

import (
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		in  string
		out string
		err error
	}{
		{"0", "zero", nil},
		{"9", "nine", nil},
		{"10", "ten", nil},
		{"11", "eleven", nil},
		{"12", "twelve", nil},
		{"13", "thirteen", nil},
		{"15", "fifteen", nil},
		{"18", "eighteen", nil},
		{"20", "twenty", nil},
		{"30", "thirty", nil},
		{"40", "forty", nil},
		{"80", "eighty", nil},
		{"100", "one hundred", nil},
		{"155", "one hundred fifty five", nil},
		{"1555", "one thousand five hundred fifty five", nil},
		{"10000", "ten thousand", nil},
		{"1234567890", "one billion two hundred thirty four million five hundred sixty seven thousand eight hundred ninety", nil},
		{"110220330440", "one hundred ten billion two hundred twenty million three hundred thirty thousand four hundred forty", nil},
		{"550660770880990", "five hundred fifty trillion six hundred sixty billion seven hundred seventy million eight hundred eighty thousand nine hundred ninety", nil},
		{"111212313515818220", "one hundred eleven quadrillion two hundred twelve trillion three hundred thirteen billion five hundred fifteen million eight hundred eighteen thousand two hundred twenty", nil},
	}

	for _, tt := range tests {
		t.Run(tt.in+" to en-us", func(t *testing.T) {
			out, err := Convert(tt.in)

			if out != tt.out {
				t.Errorf("got %s, want %s", out, tt.out)
			} else if err != tt.err {
				t.Errorf("got %v, want %v", err, tt.err)
			}
		})
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Convert(strconv.Itoa(i))
		if err != nil {
			b.Error(err)
		}
	}
}

func TestSplitNumber(t *testing.T) {
	tests := []struct {
		in     float64
		outInt decimal.Decimal
		outDec decimal.Decimal
		err    error
	}{
		{0, decimal.NewFromFloat(0), decimal.NewFromFloat(0), nil},
		{0.5, decimal.NewFromFloat(0), decimal.NewFromFloat(5), nil},
		{10, decimal.NewFromFloat(10), decimal.NewFromFloat(0), nil},
		{10.75, decimal.NewFromFloat(10), decimal.NewFromFloat(75), nil},
		{100, decimal.NewFromFloat(100), decimal.NewFromFloat(0), nil},
		{100.375, decimal.NewFromInt(100), decimal.NewFromInt(375), nil},
		{111212313515818220, decimal.NewFromInt(111212313515818220), decimal.NewFromInt(0), nil},
	}

	for _, tt := range tests {
		inDecimal := decimal.NewFromFloat(tt.in)
		testName := "split " + inDecimal.String()

		t.Run(testName, func(t *testing.T) {
			i, d, err := splitNum(inDecimal)

			if !i.Sub(tt.outInt).IsZero() {
				t.Errorf("got %v, want %v", i, tt.outInt)
			} else if !d.Sub(tt.outDec).IsZero() {
				t.Errorf("got %v, want %v", d, tt.outDec)
			} else if err != tt.err {
				t.Errorf("got %v, want %v", err, tt.err)
			}
		})
	}
}
