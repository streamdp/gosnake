package game

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type food struct {
	position []coordinate
	limit    int
}

func newFood() *food {
	return &food{
		position: []coordinate{},
		limit:    10,
	}
}

func drawFood(screen tcell.Screen, desk *desk, food *food) {
	style := tcell.StyleDefault.Background(tcell.ColorDarkMagenta)
	for i := 0; i < len(food.position); i++ {
		screen.SetContent(desk.rect.shiftX+food.position[i].x, desk.rect.shiftY+food.position[i].y, tcell.RuneCkBoard, nil, style)
	}
}

func addFood(food *food, snake *snake, desk *desk) {
	if len(food.position) < food.limit {
		freeCells := difference(&snake.body, &desk.cells)
		rand.Seed(time.Now().Unix())
		food.position = append(food.position, (*freeCells)[rand.Int()%len(*freeCells)])
	}
}
