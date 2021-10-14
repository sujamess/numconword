package options

import (
	"github.com/sujamess/numconword/currency"
	"github.com/sujamess/numconword/lang"
)

type ConvertOptions struct {
	Language     *lang.Lang
	Currency     *currency.Currency
	WithCurrency *bool
}

func MergeConvertOptions(opts ...*ConvertOptions) *ConvertOptions {
	defaultLanguage, defaultCurrency, defaultWithCurrency := lang.AmericanEnglish, currency.USD, false
	cOpts := &ConvertOptions{&defaultLanguage, &defaultCurrency, &defaultWithCurrency}

	for _, co := range opts {
		if co.Language != nil {
			cOpts.Language = co.Language
		}

		if co.Currency != nil {
			cOpts.Currency = co.Currency
		}

		if co.WithCurrency != nil {
			cOpts.WithCurrency = co.WithCurrency
		}
	}

	return cOpts
}
