package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	runewidth "github.com/mattn/go-runewidth"
	log "github.com/nuveo/logSys"
)

func say(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		w := runewidth.RuneWidth(c)
		s.SetContent(x, y, c, nil, style)
		x += w
	}
}

func main() {
	encoding.Register()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Init()
	if err != nil {
		log.Fatal(err)
	}
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.Clear()

	st := tcell.StyleDefault.Background(tcell.ColorRed)
	for {
		s.Show()
		ev := s.PollEvent()
		w, h := s.Size()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			s.SetContent(w-1, h-1, '*', nil, st)
		case *tcell.EventKey:
			s.SetContent(3, 3, ev.Rune(), nil, st)

			say(s, 0, 0, defStyle, "                                 ")
			say(s, 0, 0, defStyle, fmt.Sprintf("%x -> %v", ev.Key(), ev.Name()))

			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				os.Exit(0)
			}
		}
		<-time.After(time.Duration(100) * time.Millisecond)
	}
}
