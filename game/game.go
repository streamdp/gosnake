package game

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

type button int

//game actions
const (
	RIGHT button = 1 + iota
	LEFT
	UP
	DOWN
	QUIT
	RESTART
)

func getXY(desk *desk) coordinate {
	x := rand.Intn(desk.rect.width-3) + 2
	y := rand.Intn(desk.rect.height-2) + 1
	return coordinate{
		x: x,
		y: y,
	}
}

func newGame(screen tcell.Screen, width int, height int) (desk *desk, snake *snake, food *food) {
	desk = newDesk(newRect(screen, width, height), &deskPalette{
		inner: tcell.StyleDefault.Background(tcell.ColorBisque),
		outer: tcell.StyleDefault.Background(tcell.ColorPaleVioletRed),
	})
	snake = newSnake(getXY(desk), &snakePalette{
		head: tcell.StyleDefault.Background(tcell.Color161),
		body: tcell.StyleDefault.Background(tcell.Color170),
	})
	food = newFood()
	return
}

func restartGame(desk *desk, snake *snake, food *food) (*desk, *snake, *food) {
	desk.level = 0
	desk.score = 0
	desk.running = true
	snake.body = []coordinate{}
	snake.body = append(snake.body, getXY(desk))
	snake.length = 3
	snake.direction = getRandomDirection()
	food.position = []coordinate{}
	return desk, snake, food
}

func getEvents(screen tcell.Screen, buttonPressed chan button) {
	previousPressed := button(0)
	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEscape:
				buttonPressed <- QUIT
				return
			case tcell.KeyEnter:
				buttonPressed <- RESTART
			case tcell.KeyUp:
				if previousPressed != UP {
					buttonPressed <- UP
					previousPressed = UP
				}
			case tcell.KeyDown:
				if previousPressed != DOWN {
					buttonPressed <- DOWN
					previousPressed = DOWN
				}
			case tcell.KeyLeft:
				if previousPressed != LEFT {
					buttonPressed <- LEFT
					previousPressed = LEFT
				}
			case tcell.KeyRight:
				if previousPressed != RIGHT {
					buttonPressed <- RIGHT
					previousPressed = RIGHT
				}
			}
		}
	}
}

func validDirectionSelected(current, new button) (accept bool) {
	if current == UP && new == DOWN || current == DOWN && new == UP {
		return
	}
	if current == LEFT && new == RIGHT || current == RIGHT && new == LEFT {
		return
	}
	return true
}

// Run is the function that starts the game
func Run(width int, height int, foodLimit int) {
	var screen tcell.Screen
	var err error
	var run = true
	rand.Seed(time.Now().UnixNano())
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	if screen, err = tcell.NewScreen(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	screen.Clear()
	desk, snake, food := newGame(screen, width, height)
	food.limit = foodLimit
	keyEvents := make(chan button)
	go getEvents(screen, keyEvents)
	for run {
		select {
		case bPressed := <-keyEvents:
			switch bPressed {
			case QUIT:
				run = false
			case RESTART:
				if !desk.running {
					desk, snake, food = restartGame(desk, snake, food)
				}
			case UP, DOWN, LEFT, RIGHT:
				if validDirectionSelected(snake.direction, bPressed) {
					snake.direction = bPressed
				}
			}
		case <-time.After(time.Millisecond * 50):
		}
		drawDesk(screen, desk)
		if desk.running {
			drawSnake(screen, desk, snake)
			addFood(food, snake, desk)
			drawFood(screen, desk, food)
			ateFood(food, snake, desk)
			checkCollision(snake, desk)
		}
		screen.Show()
		time.Sleep(time.Millisecond * time.Duration(100-5*desk.level))
	}
	screen.Fini()
}
