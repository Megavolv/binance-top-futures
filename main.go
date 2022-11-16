package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/adshao/go-binance/v2"
)

func main() {
	fut_client := binance.NewFuturesClient("", "")
	info, err := fut_client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Всего торговых пар: %d\n", len(info.Symbols))

	sym_kv_data := []kv{}
	for _, symdata := range info.Symbols {
		fmt.Print(".")

		if len(filterByQuote) > 0 {
			if symdata.QuoteAsset != filterByQuote {
				continue
			}
		}

		kl := fut_client.NewKlinesService()
		res, err := kl.Limit(klines_limit).Interval(klines_interval).Symbol(symdata.Symbol).Do(context.Background())
		if err != nil {
			panic(err)
		}

		median := calcMedian(res)
		if median == 0 {
			continue
		}
		sym_kv_data = append(sym_kv_data, kv{symdata.BaseAsset, median})
	}

	fmt.Println("|")

	sort.Slice(sym_kv_data, func(i, j int) bool {
		return sym_kv_data[i].Median > sym_kv_data[j].Median
	})

	var max int
	max = symbols_limit
	if len(sym_kv_data) < symbols_limit {
		max = len(sym_kv_data)
	}

	p1, _ := calcPercentile(sym_kv_data, 97)
	p2, _ := calcPercentile(sym_kv_data, 93)
	p3, _ := calcPercentile(sym_kv_data, 88)

	for i, k := range sym_kv_data[0:max] {
		fmt.Printf("%2d %s %f\n", i, ChooseColor(k.Median, p1, p2, p3)(k.Symbol), k.Median)
	}
}
