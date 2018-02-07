package game

import (
	"fmt"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Game struct {
	display *display.Display
}

func NewGame() *Game {

	return &Game{
		UI: ui.NewUserInterface(),
	}
}

func (g *Game) Start() {

	g.UI.Render()
	go g.Update()

	for {
		switch ev := termbox.PollEvent(); ev.Type {

		case termbox.EventResize:
			g.UI.Resize()

		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			default:
				if ev.Ch >= 97 && ev.Ch <= 122 {
					g.HandleKey(ev.Ch)
				}
			}
		}
	}
}

func (g *Game) Update() {

	for {
		time.Sleep(time.Second * 1)
	}
}

func (g *Game) HandleKey(char rune) {

	fmt.Println(char)
}
