package main

import (
	"github.com/adshao/go-binance/v2/futures"
	"github.com/montanaflynn/stats"
	"github.com/shopspring/decimal"
)

const (
	PERCENTILE1 = 95
	PERCENTILE2 = 92
	PERCENTILE3 = 84
	PERCENTILE4 = 0
)

type Pair struct {
	Symbol string
	Median float64
}

func (p *Pair) SetMedian(m float64) {
	p.Median = m
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
