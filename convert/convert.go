package convert

import (
	"github.com/shopspring/decimal"
	"github.com/sujamess/convex/lang"
	"github.com/sujamess/convex/options"
)

type Convert struct {
	Integer        decimal.Decimal
	Decimal        decimal.Decimal
	ConvertOptions *options.ConvertOptions
}

func (c *Convert) ToWords() string {
	switch *c.ConvertOptions.Language {
	case lang.AmericanEnglish:
		return getEnUsWords(c.Integer, c.Decimal, c.ConvertOptions.Currency, c.ConvertOptions.WithCurrency)
	default:
		return getEnUsWords(c.Integer, c.Decimal, c.ConvertOptions.Currency, c.ConvertOptions.WithCurrency)
	}
}
