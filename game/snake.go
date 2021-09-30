package game

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type direction int

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type snakePallete struct {
	head tcell.Style
	body tcell.Style
}

type snake struct {
	body      []coord
	direction direction
	length    int
	pallete   *snakePallete
}

func newSnake(x int, y int, p *snakePallete) (s *snake) {
	snakeBody := []coord{}
	snakeBody = append(snakeBody, coord{x: x, y: y})
	return &snake{
		body:      snakeBody,
		direction: direction(rand.Intn(3)) + 1,
		length:    len(snakeBody) + 2,
		pallete:   p,
	}
}

func ateFood(food *food, snake *snake, desk *desk) {
	for i := 0; i < len(food.position); i++ {
		if food.position[i] == snake.body[0] {
			food.position = append(food.position[:i], food.position[i+1:]...)
			snake.length++
			desk.score += 100
			if snake.length%10 == 0 {
				desk.level++
			}
		}
	}
}

func moveSnake(snake *snake) {
	bodyLength := len(snake.body)
	nSegment := snake.body[bodyLength-1]
	copy := snake.body[0]
	switch snake.direction {
	case LEFT:
		snake.body[0].x -= 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, nSegment)
			}
			copy, snake.body[i] = snake.body[i], copy
		}
	case RIGHT:
		snake.body[0].x += 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, nSegment)
			}
			copy, snake.body[i] = snake.body[i], copy
		}
	case UP:
		snake.body[0].y -= 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, nSegment)
			}
			copy, snake.body[i] = snake.body[i], copy
		}
	case DOWN:
		snake.body[0].y += 1
		for i := 1; i < snake.length; i++ {
			if i >= bodyLength {
				snake.body = append(snake.body, nSegment)
			}
			copy, snake.body[i] = snake.body[i], copy
		}
	}
}

func checkCollison(s *snake, d *desk) {
	for i := 1; i < len(s.body); i++ {
		if s.body[0] == s.body[i] {
			d.running = false
		}
	}
	if (s.body[0].x == 1 || s.body[0].x == d.rect.width-2) || (s.body[0].y == 0 || s.body[0].y == d.rect.heigth-1) {
		d.running = false
	}
}

func difference(snake *[]coord, desk *[]coord) (cells *[]coord) {
	var freeCells []coord
	m := map[coord]int{}
	for _, snakeVal := range *snake {
		m[snakeVal] = 1
	}
	for _, deskVal := range *desk {
		m[deskVal] = m[deskVal] + 1
	}
	for mKey, mVal := range m {
		if mVal == 1 {
			freeCells = append(freeCells, mKey)
		}
	}
	return &freeCells
}

func drawSnake(s tcell.Screen, desk *desk, snake *snake) {
	moveSnake(snake)
	s.SetContent(desk.rect.shiftX+snake.body[0].x, desk.rect.shiftY+snake.body[0].y, tcell.RuneCkBoard, nil, snake.pallete.head)
	for i := 1; i < snake.length; i++ {
		s.SetContent(desk.rect.shiftX+snake.body[i].x, desk.rect.shiftY+snake.body[i].y, tcell.RuneBoard, nil, snake.pallete.body)
	}
	s.Show()
}
