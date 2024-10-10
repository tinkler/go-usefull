package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func ParseEth(bd string, decimals uint32, f int32) string {
	v, err := decimal.NewFromString(bd)
	if err != nil {
		panic(err)
	}
	return v.Div(decimal.NewFromBigInt(big.NewInt(1), int32(decimals))).StringFixed(f)
}
