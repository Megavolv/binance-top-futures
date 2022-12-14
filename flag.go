package main

import flag "github.com/spf13/pflag"

var filterByQuote string
var top int
var klines_limit int
var klines_interval string
var sleep int
func init() {
	flag.StringVar(&filterByQuote, "filter-by-qote", "USDT", "Filter by quote asset")
	flag.IntVar(&top, "top", 30, "Maximum number of displayed symbols")
	flag.IntVar(&klines_limit, "klines_limit", 31, "Set limit for klines")
	flag.IntVar(&sleep, "sleep", 0, "Set sleep time")
	flag.StringVar(&klines_interval, "klines_interval", "1m", "Set klines interval")
	flag.Parse()
}
