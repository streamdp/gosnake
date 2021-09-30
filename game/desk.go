package game

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

type coord struct {
	x int
	y int
}

type desk struct {
	rect    *rect
	pallete *deskPallete
	cells   []coord
	score   int
	level   int
	running bool
}

type deskPallete struct {
	outer tcell.Style
	inner tcell.Style
}

type rect struct {
	width  int
	heigth int
	shiftX int
	shiftY int
}

func drawStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		width := runewidth.RuneWidth(c)
		if width == 0 {
			comb = []rune{c}
			c = ' '
			width = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += width
	}
}

func drawDesk(s tcell.Screen, d *desk) {
	s.Clear()
	for row := 0; row < d.rect.heigth; row++ {
		for col := 0; col < d.rect.width; col++ {
			if (row == 0 || row == d.rect.heigth-1) || (col < 2 || col > d.rect.width-3) {
				s.SetContent(d.rect.shiftX+col, d.rect.shiftY+row, rune(0), nil, d.pallete.outer)
			} else {
				s.SetContent(d.rect.shiftX+col, d.rect.shiftY+row, rune(0), nil, d.pallete.inner)
			}
		}
	}
	if d.running {
		style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
		text := "Score: " + strconv.Itoa(d.score) + "  Level: " + strconv.Itoa(d.level)
		drawStr(s, d.rect.shiftX+1, 0, style, text)
	} else {
		style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkRed)
		text := "GAME OVER! YOU SCORE: " + strconv.Itoa(d.score)
		drawStr(s, d.rect.shiftX+d.rect.width/2-len([]rune(text))/2, d.rect.heigth/2, style, text)
		text = "PRESS ESC TO QUIT OR ENTER TO PLAY AGAIN"
		drawStr(s, d.rect.shiftX+d.rect.width/2-len([]rune(text))/2, d.rect.heigth/2+1, style.Reverse(true), text)
	}
	s.Show()
}

func newDesk(rect *rect, pallete *deskPallete) (readyDesk *desk) {
	var c []coord
	for i := 2; i < rect.width-2; i++ {
		for j := 1; j < rect.heigth-1; j++ {
			c = append(c, coord{x: i, y: j})
		}
	}
	return &desk{
		rect:    rect,
		pallete: pallete,
		cells:   c,
		level:   1,
		score:   0,
		running: true,
	}
}

func newRect(s tcell.Screen, width int, heigth int) (rectangle *rect) {
	windowWidth, _ := s.Size()
	return &rect{
		width:  width,
		heigth: heigth,
		shiftX: windowWidth/2 - width/2,
		shiftY: 1,
	}
}
