package game

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type food struct {
	position map[coordinate]struct{}
	limit    int
}

func newFood() *food {
	return &food{
		position: make(map[coordinate]struct{}, 10),
		limit:    10,
	}
}

func drawFood(screen tcell.Screen, desk *desk, food *food) {
	style := tcell.StyleDefault.Background(tcell.ColorDarkMagenta)
	for fp := range food.position {
		screen.SetContent(desk.rect.shiftX+fp.x, desk.rect.shiftY+fp.y, tcell.RuneCkBoard, nil, style)
	}
}

func getRandPoint(desk *desk) coordinate {
	return coordinate{
		x: 2 + rand.Intn(desk.rect.width-4),
		y: 1 + rand.Intn(desk.rect.height-2),
	}
}

func addFood(food *food, snake *snake, desk *desk) {
	if len(food.position) < food.limit {
		var newFoodPoint = coordinate{}
	loop:
		for {
			newFoodPoint = getRandPoint(desk)
			for i := range snake.body {
				_, ok := food.position[newFoodPoint]
				if ok || snake.body[i] == newFoodPoint {
					continue loop
				}
			}
			break
		}
		food.position[newFoodPoint] = struct{}{}
	}
}
