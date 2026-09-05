package game

import (
	"math/rand/v2"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

const (
	borderSize = 2
	minWidth   = 32
	minHeight  = 6
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
	inner: tcell.StyleDefault.Background(color.Bisque),
	outer: tcell.StyleDefault.Background(color.PaleVioletRed),
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
		x: borderSize + rand.IntN(d.rect.width-2*borderSize),
		y: borderSize + rand.IntN(d.rect.height-2*borderSize-1),
	}
}
