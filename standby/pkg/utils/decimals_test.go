package utils

import (
	"math/big"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

func TestParseEth(t *testing.T) {
	a := ParseEth(strings.ReplaceAll("1544012652165711426344390", " ", ""), 18, 8)
	t.Log(a)
	b := ParseEth(strings.ReplaceAll("  1930000000000000000000000000000", " ", ""), 30, 8)
	t.Log(b)
	c := ParseEth(strings.ReplaceAll("        1000000000000000000000000", " ", ""), 30, 8)
	t.Log(c)
	c1, _ := decimal.NewFromString("2000000000000000")
	c2, _ := decimal.NewFromString("63786977603970000")
	r12 := c1.Mul(c2).Div(decimal.NewFromBigInt(big.NewInt(1), 30))
	t.Log(r12.StringFixed(2))
}
