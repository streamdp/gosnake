package game

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type snakePalette struct {
	head tcell.Style
	body tcell.Style
}

type snake struct {
	body      []coordinate
	direction button
	length    int
	palette   *snakePalette
}

func getRandomDirection() button {
	return button(rand.Intn(4) + 1)
}

func newSnake(xy coordinate, palette *snakePalette) *snake {
	snakeBody := []coordinate{xy}
	return &snake{
		body:      snakeBody,
		direction: getRandomDirection(),
		length:    3,
		palette:   palette,
	}
}

func ateFood(food *food, snake *snake, desk *desk) {
	for fp := range food.position {
		if fp.x == snake.body[0].x && fp.y == snake.body[0].y {
			delete(food.position, fp)
			snake.length += 1
			desk.score += 100
			if snake.length%10 == 0 {
				desk.level += 1
			}
		}
	}
}

func moveSnake(snake *snake) {
	bodyLength := len(snake.body)
	lastSegment := snake.body[bodyLength-1]
	firstSegment := snake.body[0]
	switch snake.direction {
	case LEFT:
		snake.body[0].x -= 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, lastSegment)
			}
			firstSegment, snake.body[i] = snake.body[i], firstSegment
		}
	case RIGHT:
		snake.body[0].x += 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, lastSegment)
			}
			firstSegment, snake.body[i] = snake.body[i], firstSegment
		}
	case UP:
		snake.body[0].y -= 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, lastSegment)
			}
			firstSegment, snake.body[i] = snake.body[i], firstSegment
		}
	case DOWN:
		snake.body[0].y += 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, lastSegment)
			}
			firstSegment, snake.body[i] = snake.body[i], firstSegment
		}
	}
}

func checkCollision(snake *snake, desk *desk) {
	for i := 1; i < len(snake.body); i++ {
		if snake.body[0] == snake.body[i] {
			desk.running = false
		}
	}
	if (snake.body[0].x == 1 || snake.body[0].x == desk.rect.width-2) || (snake.body[0].y == 0 || snake.body[0].y == desk.rect.height-1) {
		desk.running = false
	}
}

func drawSnake(screen tcell.Screen, desk *desk, snake *snake) {
	moveSnake(snake)
	screen.SetContent(desk.rect.shiftX+snake.body[0].x, desk.rect.shiftY+snake.body[0].y, tcell.RuneCkBoard, nil, snake.palette.head)
	for i := 1; i < snake.length; i++ {
		screen.SetContent(desk.rect.shiftX+snake.body[i].x, desk.rect.shiftY+snake.body[i].y, tcell.RuneBoard, nil, snake.palette.body)
	}
}
