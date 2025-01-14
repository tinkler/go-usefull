package utils

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func dSalted(targetSuffix string) {
	salt := make([]byte, 32)
	var address common.Address
	i := 0
	fmt.Print("times:")
	for !strings.Contains(strings.ToLower(address.Hex()), targetSuffix) {
		_, err := rand.Read(salt)
		if err != nil {
			panic(err)
		}

		deployerAddress := common.HexToAddress("0x1ba0e5a979efce377f890e263479903c67377f41")
		initCode := []byte("abc")

		hash := crypto.Keccak256Hash(
			append(
				append(
					append([]byte{0xff, 0xff},
						deployerAddress.Bytes()...,
					),
					salt...,
				),
				crypto.Keccak256Hash(initCode).Bytes()...,
			),
		)
		address = common.BytesToAddress(hash.Bytes()[12:])
		i++
		// fmt.Printf("%d\r", i)
	}

	fmt.Printf("Salt: %x\n", salt)
	fmt.Printf("Predicted Contract Address: %s\n", address.Hex())
}
