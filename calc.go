package main

import (
	"github.com/adshao/go-binance/v2/futures"
	"github.com/montanaflynn/stats"
	"github.com/shopspring/decimal"
)

type kv struct {
	Symbol string
	Median float64
}

func calcPercentile(data []kv, percent float64) (float64, error) {

	fdata := make([]float64, len(data))

	for i, k := range data {
		fdata[i] = k.Median
	}

	return stats.Percentile(fdata, percent)
}

func calcMedian(data []*futures.Kline) (median float64) {

	floatdata := make([]float64, len(data)-1)
	for i, kline := range data[0 : len(data)-1] {

		h, _ := decimal.NewFromString(kline.High)
		l, _ := decimal.NewFromString(kline.Low)

		perc := h.Sub(l).Div(h.Add(l).Div(decimal.NewFromInt(2))).Mul(decimal.NewFromInt(100))
		floatdata[i], _ = perc.Float64()
	}

	median, _ = stats.Median(floatdata)

	return
}
