package game

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

type coordinate struct {
	x int
	y int
}

type desk struct {
	rect    *rect
	palette *deskPalette
	score   int
	level   int
	running bool
}

type deskPalette struct {
	outer tcell.Style
	inner tcell.Style
}

type rect struct {
	width  int
	height int
	shiftX int
	shiftY int
}

func drawStr(screen tcell.Screen, x int, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		width := runewidth.RuneWidth(c)
		if width == 0 {
			comb = []rune{c}
			c = ' '
			width = 1
		}
		screen.SetContent(x, y, c, comb, style)
		x += width
	}
}

func drawDesk(screen tcell.Screen, desk *desk) {
	for row := 0; row < desk.rect.height; row++ {
		for col := 0; col < desk.rect.width; col++ {
			if (row == 0 || row == desk.rect.height-1) || (col < 2 || col > desk.rect.width-3) {
				screen.SetContent(desk.rect.shiftX+col, desk.rect.shiftY+row, tcell.RuneCkBoard, nil, desk.palette.outer)
			} else {
				screen.SetContent(desk.rect.shiftX+col, desk.rect.shiftY+row, rune(' '), nil, desk.palette.inner)
			}
		}
	}
	if desk.running {
		style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
		text := "Score: " + strconv.Itoa(desk.score) + "  Level: " + strconv.Itoa(desk.level)
		drawStr(screen, desk.rect.shiftX+1, 0, style, text)
	} else {
		style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkRed)
		text := "GAME OVER! YOU SCORE: " + strconv.Itoa(desk.score)
		drawStr(screen, desk.rect.shiftX+desk.rect.width/2-len([]rune(text))/2, desk.rect.height/2, style, text)
		text = "PRESS ESC TO QUIT OR ENTER TO PLAY AGAIN"
		drawStr(screen, desk.rect.shiftX+desk.rect.width/2-len([]rune(text))/2, desk.rect.height/2+1, style.Reverse(true), text)
	}
}

func newDesk(rect *rect, palette *deskPalette) *desk {
	return &desk{
		rect:    rect,
		palette: palette,
		level:   1,
		score:   0,
		running: true,
	}
}

func newRect(screen tcell.Screen, width int, height int) *rect {
	sWidth, _ := screen.Size()
	return &rect{
		width:  width,
		height: height,
		shiftX: sWidth/2 - width/2,
		shiftY: 1,
	}
}
