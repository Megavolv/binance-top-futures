package main

import "github.com/fatih/color"

var yellow func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()
var red func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()
var green func(a ...interface{}) string = color.New(color.FgGreen).SprintFunc()
var blue func(a ...interface{}) string = color.New(color.FgBlue).SprintFunc()
var cyan func(a ...interface{}) string = color.New(color.FgCyan).SprintFunc()
var magenta func(a ...interface{}) string = color.New(color.FgMagenta).SprintFunc()

func (l *List) chooseColor(p *Pair) func(a ...interface{}) string {

	if p.Median >= l.p1 {
		return red
	} else if p.Median >= l.p2 {
		return yellow
	} else if p.Median >= l.p3 {
		return cyan
	}

	return blue
}
