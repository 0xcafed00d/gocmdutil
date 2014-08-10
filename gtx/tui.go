package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
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


func termtest() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			printAt(0, 1, fmt.Sprint(ev), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()

			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
			case termbox.KeyArrowDown:
			}

			termbox.Flush()

		case termbox.EventResize:
			x, y := ev.Width, ev.Height
			printAt(0, 0, fmt.Sprintf("[%d, %d] ", x, y), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
		}
	}

	termbox.Flush()
}

