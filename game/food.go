package game

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type food struct {
	position []coord
	limit    int
}

func newFood() (f *food) {
	return &food{
		position: []coord{},
		limit:    10,
	}
}

func drawFood(s tcell.Screen, d *desk, f *food) {
	style := tcell.StyleDefault.Background(tcell.ColorDarkMagenta)
	for i := 0; i < len(f.position); i++ {
		s.SetContent(d.rect.shiftX+f.position[i].x, d.rect.shiftY+f.position[i].y, tcell.RuneCkBoard, nil, style)
	}
	s.Show()
}

func addFood(f *food, s *snake, d *desk) {
	if len(f.position) < f.limit {
		freeCells := difference(&s.body, &d.cells)
		rand.Seed(time.Now().Unix())
		f.position = append(f.position, (*freeCells)[rand.Int()%len(*freeCells)])
	}
}
