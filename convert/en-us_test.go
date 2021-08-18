package convert

import (
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestToEnUs(t *testing.T) {
	tests := []struct {
		in  int64
		out string
	}{
		{0, "zero"},
		{9, "nine"},
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{15, "fifteen"},
		{18, "eighteen"},
		{20, "twenty"},
		{30, "thirty"},
		{40, "forty"},
		{80, "eighty"},
		{100, "one hundred"},
		{155, "one hundred fifty five"},
		{1555, "one thousand five hundred fifty five"},
		{10000, "ten thousand"},
		{1234567890, "one billion two hundred thirty four million five hundred sixty seven thousand eight hundred ninety"},
		{110220330440, "one hundred ten billion two hundred twenty million three hundred thirty thousand four hundred forty"},
		{550660770880990, "five hundred fifty trillion six hundred sixty billion seven hundred seventy million eight hundred eighty thousand nine hundred ninety"},
		{111212313515818220, "one hundred eleven quadrillion two hundred twelve trillion three hundred thirteen billion five hundred fifteen million eight hundred eighteen thousand two hundred twenty"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.in, 10)+" to en-us", func(t *testing.T) {
			out := toEnUs(decimal.NewFromInt(tt.in))

			if out != tt.out {
				t.Errorf("got %s, want %s", out, tt.out)
			}
		})
	}
}
