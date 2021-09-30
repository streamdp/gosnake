package game

import (
	"reflect"
	"testing"

	"github.com/gdamore/tcell"
)

func TestGetXY(t *testing.T) {
	desk := &desk{
		rect: &rect{
			width:  70,
			heigth: 20,
			shiftX: 10,
			shiftY: 1,
		},
		palette: &deskPalette{
			outer: tcell.StyleDefault,
			inner: tcell.StyleDefault,
		},
		cells: []coord{
			{
				x: 1,
				y: 1,
			},
		},
		score:   0,
		level:   1,
		running: true,
	}
	got := getXY(desk)

	if !IsInstanceOf(got, coord{}) {
		t.Errorf("got %v; should be %v", reflect.TypeOf(got), reflect.TypeOf(coord{}))
	}
	if got.x < 2 || got.x > desk.rect.width-3 {
		t.Errorf("got %d; should be [%v..%v]", got, 2, desk.rect.width-3)
	}
	if got.y < 1 || got.y > desk.rect.heigth-2 {
		t.Errorf("got %d; should be [%v..%v]", got, 1, desk.rect.heigth-2)
	}
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}
