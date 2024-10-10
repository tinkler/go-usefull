package main

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type PointData struct {
	Account       string          `json:"account"`
	Liquidity     decimal.Decimal `json:"liquidity"`
	Referral      decimal.Decimal `json:"referral"`
	TraderProfit  decimal.Decimal `json:"traderProfit"`
	TradingVolume decimal.Decimal `json:"tradingVolume"`
	Swap          decimal.Decimal `json:"swap"`
	Total         decimal.Decimal `json:"total"`
}

func parseFromCSV() ([]*PointData, error) {
	file, err := os.Open("point-1_2024-06-30.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Header", header)
	var data []*PointData
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// account
		if !common.IsHexAddress(record[0]) {
			continue
		}
		var (
			lp, referral, profit, trading, swap, total decimal.Decimal
		)
		if len(record[1]) > 0 {
			var err error
			lp, err = decimal.NewFromString(record[1])
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(record[2]) > 0 {
			var err error
			referral, err = decimal.NewFromString(record[2])
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(record[3]) > 0 {
			var err error
			profit, err = decimal.NewFromString(record[3])
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(record[4]) > 0 {
			var err error
			trading, err = decimal.NewFromString(record[4])
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(record[5]) > 0 {
			var err error
			swap, err = decimal.NewFromString(record[5])
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(record[6]) > 0 {
			var err error
			total, err = decimal.NewFromString(record[6])
			if err != nil {
				log.Fatal(err)
			}

		}

		d := PointData{
			Account:       record[0],
			Liquidity:     lp,
			Referral:      referral,
			TraderProfit:  profit,
			TradingVolume: trading,
			Swap:          swap,
			Total:         total,
		}
		data = append(data, &d)
	}
	return data, nil
}

func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func writeToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 644)
}

func main() {
	data, err := parseFromCSV()
	if err != nil {
		log.Fatal(err)
	}
	byt, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	compressedData, err := compress(byt)
	if err != nil {
		log.Fatal(err)
	}
	if err := writeToFile("point-1.gz", compressedData); err != nil {
		log.Fatal(err)
	}
}
