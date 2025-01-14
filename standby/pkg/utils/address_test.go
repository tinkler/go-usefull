package utils

import (
	"strings"
	"testing"
)

func TestToLower(t *testing.T) {
	println(strings.ToLower("0x03F0ba70Fc5775a824E7662c50AC02BBDE84e1aE"))
	println(strings.ToLower("0x626aba32e69Df9a2d4A4Db1ad7E9C74c513326Ca"))
	println(strings.ToLower("0xE2dbADe2d5dE40853A96a7Ff811a7651280486b3"))
	println(strings.ToLower("0xc2ef1173a4B051BbCc4474c407BBb2ac7E02e363"))
	println(strings.ToLower("0xC764f60FaF30bE89C3D41ef0f44e8E8bF1204A12"))
	println(strings.ToLower("0x1477f60D8A4666366aa271765ef240BceEEbe8f9"))
	println(strings.ToLower("0xdFb06147A4Bd5A3D66824a1a717Cc6A06e6aC543"))
	println(strings.ToLower("0x593851CA21Bfd716a9FC71c744e8b1CdF101535b"))
	println(strings.ToLower("0x76A5AE1eC8EB2Aeee7692077A147869aBeAc8A24"))
	println(strings.ToLower("0xdF7Eca6aF127C4F562B800B38Fbe8fFEfc7df5f0"))
	println(strings.ToLower("0x8b4C8E2bC31623Fe90D9407e697dF520a737404c"))
	println(strings.ToLower("0x6BAc80315aef7291c8aE23ff0a52c554a3f0b598"))
	println(strings.ToLower("0x57696a563B66dFA0c180747217d0EdCF17A9cD51"))
	println(strings.ToLower("0xAd3610626505b88d34A1f129b06cd18902b62cd1"))
	println(strings.ToLower("0xef15E309b26fc928C1fD36EF0fd7015bb3a80D5E"))
	println(strings.ToLower("0xba22EdEBb9738DcF29F16aC1aDbDe03A25EBF795"))
	println(strings.ToLower("0xD058E7bCa54603e46b1128078Fb92c31953cf029"))
	println(strings.ToLower("0xf79cD7b73c86cf051D6E814c4E457081e019F85e"))
	println(strings.ToLower("0x46931a27aDa075495CB0A2EB21Bd004a330D4A72"))
	println(strings.ToLower("0x96E1991eBD4b08975cBDca89d09FE7b677Af407B"))
}

func TestDSalted(t *testing.T) {
	dSalted("seed")
}
