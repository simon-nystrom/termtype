package main

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/simon-nystrom/termtype/game"
)

var (
	width, height int
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	g := game.NewGame()
	g.Start()
}

func WriteWord(x, y int, s string, textColor termbox.Attribute) {
	for _x, char := range s {
		termbox.SetCell(_x+x, y, char, textColor, termbox.ColorDefault)
	}
}

func WriteEscapeSequence() {
	esc := "Escape key to exit"
	WriteWord(width-len(esc), height-1, esc, termbox.ColorRed)
}
