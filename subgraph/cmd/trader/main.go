package main

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
	"github.com/shopspring/decimal"
	"github.com/tinkler/moonmist/pkg/mlog"
	"github.com/xuri/excelize/v2"
)

const totalSupply int64 = 10000000

var traderSupply = decimal.NewFromFloat(float64(totalSupply) * float64(0.2))

const gqlTrader = `
query MyQuery($skip: Int!) {
  userTraderAirs(skip: $skip) {
    account
    tradeFeeUSD
  }
}
`

type TraderAirUserEntity struct {
	Account     string `json:"account"`
	TradeFeeUSD string `json:"tradeFeeUSD"`
}

type traderResponse struct {
	UserTraderAirs []*TraderAirUserEntity `json:"userTraderAirs"`
}

func main() {
	mlog.ConsoleLevel = mlog.L_DEBUG
	ctx := context.Background()
	client := graphql.NewClient("https://graph.loxodrome.xyz/subgraphs/name/loxodrome/dashboard")
	req := graphql.NewRequest(gqlTrader)
	req.Header.Set("Cache-Control", "no-cache")
	req.Var("skip", 0)
	var datas []*TraderAirUserEntity
	nextPage := 0
REQUEST:
	var respData traderResponse
	if err := client.Run(ctx, req, &respData); err != nil {
		mlog.Error("can not sync lp, %v", err)
		os.Exit(0)
	}
	datas = append(datas, respData.UserTraderAirs...)
	if len(respData.UserTraderAirs) == 100 {
		nextPage++
		req.Var("skip", nextPage*100)
		goto REQUEST
	}

	total := decimal.Zero
	for _, data := range datas {
		userUSD, err := decimal.NewFromString(data.TradeFeeUSD)
		if err != nil {
			mlog.Error(err)
			os.Exit(0)
		}
		total = total.Add(userUSD)
	}
	if total.Equal(decimal.Zero) {
		mlog.Error("total is 0")
		os.Exit(0)
	}

	arr := make([][]interface{}, 0, len(datas))
	totalEstimated := decimal.Zero
	for _, data := range datas {
		userUSD, _ := decimal.NewFromString(data.TradeFeeUSD)
		estimated := userUSD.Mul(traderSupply).Div(total)
		totalEstimated = totalEstimated.Add(estimated)
		arr = append(arr, []interface{}{data.Account, data.TradeFeeUSD, estimated.String()})
	}
	mlog.Info(total.String(), totalEstimated.String())

	exportExcel(arr)
}

func exportExcel(data [][]interface{}) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.NewSheet("Trader")
	for idx, row := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println("Row overflow")
			os.Exit(0)
		}
		f.SetSheetRow("Trader", cell, &row)
	}
	if err := f.SaveAs("LoxodromeTrader.xlsx"); err != nil {
		fmt.Println(err)
	}
}
