package main

import (
	"context"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

func main() {

	tube := make(chan Pair, 16)

	go func(p chan Pair) {
		list := NewList()
		for elem := range p {
			list.Set(elem.Symbol, elem.Median)
			list.Show()
		}
	}(tube)

	fut_client := binance.NewFuturesClient("", "")
	info, err := fut_client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Всего торговых пар: %d\n", len(info.Symbols))

	const numJobs = 5
	jobs := make(chan string, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, tube, fut_client)
	}

	for {
		for _, symdata := range info.Symbols {
			if len(filterByQuote) > 0 {
				if symdata.QuoteAsset != filterByQuote {
					continue
				}
			}

			jobs <- symdata.Symbol
		}
	}
}

func worker(id int, jobs <-chan string, tube chan Pair, fut_client *futures.Client) {
	for sym := range jobs {
		kl := fut_client.NewKlinesService()
		res, err := kl.Limit(klines_limit).Interval(klines_interval).Symbol(sym).Do(context.Background())
		if err != nil {
			panic(err)
		}

		median := calcMedian(res)
		if median == 0 {
			continue
		}

		tube <- Pair{Symbol: sym, Median: median}

		time.Sleep(2 * time.Second)
	}
}
