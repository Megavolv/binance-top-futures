package main

import (
	"fmt"
	"sort"

	"github.com/montanaflynn/stats"
)

type List struct {
	Pairs          []*Pair
	p1, p2, p3, p4 float64
}

func (l *List) getPercentileBy(percent float64) (float64, error) {

	fdata := make([]float64, l.Len())

	for i, k := range l.Pairs {
		fdata[i] = k.Median
	}

	return stats.Percentile(fdata, percent)
}

func (l *List) recalcAllPercentiles() {
	l.p1, _ = l.getPercentileBy(PERCENTILE1)
	l.p2, _ = l.getPercentileBy(PERCENTILE2)
	l.p3, _ = l.getPercentileBy(PERCENTILE3)
}

func (l *List) matchP(p *Pair) int {
	if p.Median >= l.p1 {
		return PERCENTILE1
	} else if p.Median >= l.p2 {
		return PERCENTILE2
	} else if p.Median >= l.p3 {
		return PERCENTILE3
	}

	return PERCENTILE4
}

func NewList() List {
	return List{
		Pairs: []*Pair{},
	}
}

func (l *List) Set(sym string, val float64) {
	l.get(sym).SetMedian(val)
	l.Sort()
	l.recalcAllPercentiles()
}

func (l *List) get(sym string) *Pair {
	for _, s := range l.Pairs {
		if s.Symbol == sym {
			return s
		}
	}

	p := &Pair{Symbol: sym, Median: 0}

	l.add(p)
	return p
}

func (l List) Len() int {
	return len(l.Pairs)
}

func (l List) Sort() {
	sort.Slice(l.Pairs, func(i, j int) bool {
		return l.Pairs[i].Median > l.Pairs[j].Median
	})
}

func (l *List) add(p *Pair) {
	l.Pairs = append(l.Pairs, p)
}

func (l *List) Show() {
	fmt.Print("\033[H\033[2J")

	var max int
	max = top
	if l.Len() < top {
		max = l.Len()
	}

	for i, k := range l.Pairs[0:max] {
		fmt.Printf("%2d %s %f\n", i, l.chooseColor(k)(k.Symbol), k.Median)
	}
}
