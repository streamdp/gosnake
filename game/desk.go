package game

import (
	"math/rand"

	"github.com/gdamore/tcell"
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
}

type deskPalette struct {
	outer tcell.Style
	inner tcell.Style
}

var defaultDeskPalette = &deskPalette{
	inner: tcell.StyleDefault.Background(tcell.ColorBisque),
	outer: tcell.StyleDefault.Background(tcell.ColorPaleVioletRed),
}

type rect struct {
	width  int
	height int
	shiftX int
	shiftY int
}

func newDesk(rect *rect, palette *deskPalette) *desk {
	return &desk{
		rect:    rect,
		palette: palette,
		level:   1,
		score:   0,
	}
}

func newRect(screenSize, width int, height int) *rect {
	return &rect{
		width:  width,
		height: height,
		shiftX: screenSize/2 - width/2,
		shiftY: 1,
	}
}

func (d *desk) getRandPoint() coordinate {
	return coordinate{
		x: 2 + rand.Intn(d.rect.width-4),
		y: 1 + rand.Intn(d.rect.height-2),
	}
}
