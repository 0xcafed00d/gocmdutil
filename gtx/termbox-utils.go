package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/go_sandbox/geom"
	"unicode/utf8"
)

func printAt(x, y int, s string, fg, bg termbox.Attribute) {
	for len(s) > 0 {
		r, rlen := utf8.DecodeRuneInString(s)
		termbox.SetCell(x, y, r, fg, bg)
		s = s[rlen:]
		x++
	}
}

func printAtDef(x, y int, s string) {
	printAt(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func clearRect(rect geom.Rectangle, c rune, fg, bg termbox.Attribute) {
	w, h := termbox.Size()
	sz := geom.RectangleFromSize(geom.Coord{w, h})

	toClear, ok := geom.RectangleIntersection(rect, sz)
	if ok {
		for y := toClear.Min.Y; y < toClear.Max.Y; y++ {
			for x := toClear.Min.X; x < toClear.Max.X; x++ {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}

func clearRectDef(rect geom.Rectangle) {
	clearRect(rect, ' ', termbox.ColorDefault, termbox.ColorDefault)
}
