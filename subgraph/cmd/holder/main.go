package main

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
	"github.com/tinkler/moonmist/pkg/mlog"
	"github.com/xuri/excelize/v2"
)

const gqlSIDEHolder = `
query MyQuery($skip: Int!) {
  holders(skip: $skip) {
    account
    balance
    id
  }
}
`

type HolderEntity struct {
	Account string
	Balance string
}

type holderResponse struct {
	Holders []*HolderEntity `json:"holders"`
}

func main() {
	ctx := context.Background()
	client := graphql.NewClient("http://61.10.9.22:10004/subgraphs/name/we/erc-holder")
	req := graphql.NewRequest(gqlSIDEHolder)
	req.Header.Set("Cache-Control", "no-cache")
	req.Var("skip", 0)
	var datas []*HolderEntity
	nextPage := 0
REQUEST:
	var respData holderResponse
	if err := client.Run(ctx, req, &respData); err != nil {
		mlog.Error("can not sync lp, %v", err)
		os.Exit(0)
	}
	datas = append(datas, respData.Holders...)
	if len(respData.Holders) == 100 {
		nextPage++
		req.Var("skip", nextPage*100)
		goto REQUEST
	}

	arr := make([][]interface{}, 0, len(datas))
	for _, data := range datas {
		arr = append(arr, []interface{}{data.Account, data.Balance})
	}

	exportExcel(arr)
}

func exportExcel(data [][]interface{}) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.NewSheet("Holder")
	for idx, row := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println("Row overflow")
			os.Exit(0)
		}
		f.SetSheetRow("Holder", cell, &row)
	}
	if err := f.SaveAs("SIDEHolder.xlsx"); err != nil {
		fmt.Println(err)
	}
}
