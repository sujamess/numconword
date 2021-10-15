package options

import (
	"github.com/sujamess/numconword/currency"
	"github.com/sujamess/numconword/lang"
)

// ConvertOptions represents options that can be used to configure a cv operation.
type ConvertOptions struct {
	// A language for selecting a language for words. Default is lang.AmericanEnglish
	Language *lang.Lang

	// A currency for selecting a currency to add to words. Default is currency.USD
	Currency *currency.Currency

	// If true, the returned readable words will have a currency
	WithCurrency *bool
}

// MergeConvertOptions combines the given ConvertOptions instances into a single ConvertOptions in a last-one-wins fashion.
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
