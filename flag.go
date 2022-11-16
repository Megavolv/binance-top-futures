package main

import flag "github.com/spf13/pflag"

var filterByQuote string
var symbols_limit int
var klines_limit int
var klines_interval string

func init() {
	flag.StringVar(&filterByQuote, "filter-by-qote", "USDT", "Filter by quote asset")
	flag.IntVar(&symbols_limit, "symbols_limit", 20, "Maximum number of displayed symbols")
	flag.IntVar(&klines_limit, "klines_limit", 15, "Set limit for klines")
	flag.StringVar(&klines_interval, "klines_interval", "1m", "Set klines interval")
	flag.Parse()
}