package convert

import (
	"github.com/shopspring/decimal"

	"github.com/sujamess/numconword/currency"
)

func getEnUsWords(i, d decimal.Decimal, c *currency.Currency, withCurrency *bool) string {
	minusText, integerText, decimalText := "minus", toEnUs(i.Abs()), toEnUs(d)

	var res string

	if i.IsNegative() {
		res += minusText
	}

	res += integerText + " "

	if withCurrency != nil && *withCurrency {
		res += c.String() + " "
	}

	if !d.IsZero() && d.IsPositive() {
		res += currency.Point + " " + decimalText + " "
	}

	if withCurrency != nil && *withCurrency {
		res += currency.CoinCurrencies[*c].String() + " "

		return res
	}

	return res[:len(res)-1]
}

func toEnUs(n decimal.Decimal) string {
	units := map[int64]string{
		2:  "hundred",
		3:  "thousand",
		6:  "million",
		9:  "billion",
		12: "trillion",
		15: "quadrillion",
		18: "quintillion",
		21: "sextillion",
		24: "septillion",
		27: "octillion",
		30: "nonillion",
	}

	hash := map[int64]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}

	if v, exist := hash[n.IntPart()]; exist {
		return v
	}

	var res string

	for l := len(n.String()); n.IsPositive(); l-- {
		if l == 2 {
			h := n.Div(decimal.NewFromInt(10))
			u := hash[decimal.NewFromInt(h.IntPart()).Mul(decimal.NewFromInt(10)).IntPart()]

			res += u + " "

			if n.Sub(decimal.NewFromInt(h.IntPart()).Mul(decimal.NewFromInt(10))).IsPositive() {
				lastWord := hash[n.Sub(decimal.NewFromInt(h.IntPart()).Mul(decimal.NewFromInt(10))).IntPart()] + " "

				if len(lastWord) > 0 {
					res += lastWord
				}
			}

			n = n.Sub(h.Mul(decimal.NewFromInt(10)).Add(n.Div(decimal.NewFromInt(10))))
		} else {
			unitText, exist := units[int64(l-1)]
			if !exist {
				continue
			}

			digitLeft := n.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(l - 1)))).IntPart()

			if h, exist := hash[digitLeft]; exist {
				res += h + " " + unitText + " "
			} else {
				dl := digitLeft

				if dl/100 > 0 {
					dh := dl / 100
					u := units[2]

					res += hash[dh] + " " + u + " "
					dl -= dh * 100
				}

				if dhv, exist := hash[dl]; exist {
					res += dhv + " " + unitText + " "
				} else {
					dh := dl / 10
					du := hash[dh*10]

					var lastWord string

					if dl-dh*10 > 0 {
						lastWord = hash[dl-dh*10] + " "
					}

					res += du + " " + lastWord + unitText + " "
				}
			}

			n = n.Sub(decimal.NewFromInt(digitLeft).Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(l - 1)))))
		}
	}

	return res[:len(res)-1]
}
