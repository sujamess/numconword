package convert

import (
	"github.com/shopspring/decimal"
	"github.com/sujamess/numconword/lang"
	"github.com/sujamess/numconword/options"
)

// Convert represents fields that used to get words
type Convert struct {
	// An integer value of the input
	Integer decimal.Decimal

	// A decimal value of the input
	Decimal decimal.Decimal

	// An options for configure words
	ConvertOptions *options.ConvertOptions
}

// ToWords converts the input to words with options.ConvertOptions
func (c *Convert) ToWords() string {
	switch *c.ConvertOptions.Language {
	case lang.AmericanEnglish:
		return getEnUsWords(c.Integer, c.Decimal, c.ConvertOptions.Currency, c.ConvertOptions.WithCurrency)
	default:
		return getEnUsWords(c.Integer, c.Decimal, c.ConvertOptions.Currency, c.ConvertOptions.WithCurrency)
	}
}
