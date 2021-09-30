package game

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type snakePallete struct {
	head tcell.Style
	body tcell.Style
}

type snake struct {
	body      []coord
	direction button
	length    int
	pallete   *snakePallete
}

func getRandomDirection() (bttn button) {
	return button(rand.Intn(3) + 1)
}

func newSnake(xy coord, pallete *snakePallete) (s *snake) {
	snakeBody := []coord{xy}
	return &snake{
		body:      snakeBody,
		direction: getRandomDirection(),
		length:    3,
		pallete:   pallete,
	}
}

func ateFood(food *food, snake *snake, desk *desk) {
	for i := 0; i < len(food.position); i++ {
		if food.position[i] == snake.body[0] {
			food.position = append(food.position[:i], food.position[i+1:]...)
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

func checkCollison(snake *snake, desk *desk) {
	for i := 1; i < len(snake.body); i++ {
		if snake.body[0] == snake.body[i] {
			desk.running = false
		}
	}
	if (snake.body[0].x == 1 || snake.body[0].x == desk.rect.width-2) || (snake.body[0].y == 0 || snake.body[0].y == desk.rect.heigth-1) {
		desk.running = false
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

func drawSnake(screen tcell.Screen, desk *desk, snake *snake) {
	moveSnake(snake)
	screen.SetContent(desk.rect.shiftX+snake.body[0].x, desk.rect.shiftY+snake.body[0].y, tcell.RuneCkBoard, nil, snake.pallete.head)
	for i := 1; i < snake.length; i++ {
		screen.SetContent(desk.rect.shiftX+snake.body[i].x, desk.rect.shiftY+snake.body[i].y, tcell.RuneBoard, nil, snake.pallete.body)
	}
	screen.Show()
}
