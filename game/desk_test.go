package game

import (
	"reflect"
	"testing"

	"github.com/gdamore/tcell"
)

func TestGetXY(t *testing.T) {
	d := &desk{
		rect: &rect{
			width:  70,
			height: 20,
			shiftX: 10,
			shiftY: 1,
		},
		palette: &deskPalette{
			outer: tcell.StyleDefault,
			inner: tcell.StyleDefault,
		},
		score: 0,
		level: 1,
	}
	got := d.getRandPoint()

	if !IsInstanceOf(got, coordinate{}) {
		t.Errorf("got %v; should be %v", reflect.TypeOf(got), reflect.TypeOf(coordinate{}))
	}
	if got.x < 2 || got.x > d.rect.width-3 {
		t.Errorf("got %d; should be [%v..%v]", got, 2, d.rect.width-3)
	}
	if got.y < 1 || got.y > d.rect.height-2 {
		t.Errorf("got %d; should be [%v..%v]", got, 1, d.rect.height-2)
	}
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}
