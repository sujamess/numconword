package numconword

import (
	"errors"

	"github.com/shopspring/decimal"

	"github.com/sujamess/numconword/convert"

	"github.com/sujamess/numconword/options"
)

var (
	ErrInvalidNumber = errors.New("convex: invalid number")
)

func ConvertString(num string, opts ...*options.ConvertOptions) (string, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return "", ErrInvalidNumber
	}

	return cv(dec, opts...)
}

func ConvertInt(num int, opts ...*options.ConvertOptions) (string, error) {
	return cv(decimal.NewFromInt(int64(num)), opts...)
}

func ConvertInt32(num int32, opts ...*options.ConvertOptions) (string, error) {
	return cv(decimal.NewFromInt32(num), opts...)
}

func ConvertInt64(num int64, opts ...*options.ConvertOptions) (string, error) {
	return cv(decimal.NewFromInt(num), opts...)
}

func ConvertFloat32(num float32, opts ...*options.ConvertOptions) (string, error) {
	return cv(decimal.NewFromFloat32(num), opts...)
}

func ConvertFloat64(num float64, opts ...*options.ConvertOptions) (string, error) {
	return cv(decimal.NewFromFloat(num), opts...)
}

func cv(dec decimal.Decimal, opts ...*options.ConvertOptions) (string, error) {
	i, d := splitNum(dec)

	converter := &convert.Convert{Integer: i, Decimal: d, ConvertOptions: options.MergeConvertOptions(opts...)}

	return converter.ToWords(), nil
}

func splitNum(n decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	integer := decimal.NewFromInt(n.IntPart())
	floating := n.Sub(integer)
	floating = floating.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(len(floating.String()) - 2))))

	return integer, floating
}
