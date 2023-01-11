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

var defaultSnakePalette = &snakePalette{
	head: tcell.StyleDefault.Background(tcell.Color161),
	body: tcell.StyleDefault.Background(tcell.Color170),
}

func getRandomDirection() button {
	return button(rand.Intn(4) + 1)
}

func newSnake(xy coordinate, palette *snakePalette) *snake {
	return &snake{
		body:      []coordinate{xy},
		direction: getRandomDirection(),
		length:    3,
		palette:   palette,
	}
}

func (s *snake) ateFood(food *food, desk *desk) {
	for fp := range food.position {
		if fp.x == s.body[0].x && fp.y == s.body[0].y {
			delete(food.position, fp)
			s.length += 1
			desk.score += 100
			if s.length%10 == 0 {
				desk.level += 1
			}
		}
	}
}

func (s *snake) setDirection(new button) {
	if s.direction == UP && new == DOWN || s.direction == DOWN && new == UP {
		return
	}
	if s.direction == LEFT && new == RIGHT || s.direction == RIGHT && new == LEFT {
		return
	}
	s.direction = new
	return
}

func (s *snake) moveSnake() {
	bodyLength := len(s.body)
	lastSegment := s.body[bodyLength-1]
	firstSegment := s.body[0]
	switch s.direction {
	case LEFT:
		s.body[0].x -= 1
		for i := 1; i < s.length; i++ {
			if i >= bodyLength {
				s.body = append(s.body, lastSegment)
			}
			firstSegment, s.body[i] = s.body[i], firstSegment
		}
	case RIGHT:
		s.body[0].x += 1
		for i := 1; i < s.length; i++ {
			if i >= bodyLength {
				s.body = append(s.body, lastSegment)
			}
			firstSegment, s.body[i] = s.body[i], firstSegment
		}
	case UP:
		s.body[0].y -= 1
		for i := 1; i < s.length; i++ {
			if i >= bodyLength {
				s.body = append(s.body, lastSegment)
			}
			firstSegment, s.body[i] = s.body[i], firstSegment
		}
	case DOWN:
		s.body[0].y += 1
		for i := 1; i < s.length; i++ {
			if i >= bodyLength {
				s.body = append(s.body, lastSegment)
			}
			firstSegment, s.body[i] = s.body[i], firstSegment
		}
	}
}
