package loxodrome

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestMinterV2(t *testing.T) {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "https://babel-api.testnet.iotex.one")
	assert.NoError(t, err)
	minterV2, err := NewMinterV2Caller(common.HexToAddress("0x7d4368B1E757157eBF332bb00CB645Fbb0DBB505"), client)
	assert.NoError(t, err)
	weeklyEmission, err := minterV2.Weekly(&bind.CallOpts{
		From: common.HexToAddress("0x39ed63a65AD05e623d641669c336769b51eDEF8C"),
	})
	assert.NoError(t, err)
	if weeklyEmission.Cmp(big.NewInt(0)) < 0 {
		t.Fail()
	} else {
		t.Log(weeklyEmission.String())
	}
}
