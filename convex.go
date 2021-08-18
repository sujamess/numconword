package convex

import (
	"errors"

	"github.com/shopspring/decimal"

	"github.com/sujamess/convex/convert"

	"github.com/sujamess/convex/options"
)

var (
	ErrInvalidNumber = errors.New("convex: invalid number")
)

func Convert(n string, opts ...*options.ConvertOptions) (string, error) {
	dec, err := decimal.NewFromString(n)
	if err != nil {
		return "", ErrInvalidNumber
	}

	i, d, err := splitNum(dec)
	if err != nil {
		return "", ErrInvalidNumber
	}

	converter := &convert.Convert{Integer: i, Decimal: d, ConvertOptions: options.MergeConvertOptions(opts...)}

	return converter.ToWords(), nil
}

func splitNum(n decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	integer := decimal.NewFromInt(n.IntPart())
	floating := n.Sub(integer)
	floating = floating.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(len(floating.String()) - 2))))

	return integer, floating, nil
}
