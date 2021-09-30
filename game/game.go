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

func getXY(desk *desk) (xy coord) {
	x := rand.Intn(desk.rect.width-3) + 2
	y := rand.Intn(desk.rect.heigth-2) + 1
	return coord{
		x: x,
		y: y,
	}
}

func newGame(screen tcell.Screen, width int, heigth int) (desk *desk, snake *snake, food *food) {
	desk = newDesk(newRect(screen, width, heigth), &deskPalette{
		inner: tcell.StyleDefault.Background(tcell.ColorBisque),
		outer: tcell.StyleDefault.Background(tcell.ColorPaleVioletRed),
	})
	snake = newSnake(getXY(desk), &snakePalette{
		head: tcell.StyleDefault.Background(tcell.Color161),
		body: tcell.StyleDefault.Background(tcell.Color170),
	})
	food = newFood()
	return desk, snake, food
}

func restartGame(screen tcell.Screen, desk *desk, snake *snake, food *food) (d *desk, s *snake, f *food) {
	desk.level = 0
	desk.score = 0
	desk.running = true
	snake.body = []coord{}
	snake.body = append(snake.body, getXY(desk))
	snake.length = 3
	snake.direction = getRandomDirection()
	food.position = []coord{}
	return desk, snake, food
}

func getEvents(screen tcell.Screen, buttonPressed chan button) {
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
				buttonPressed <- UP
			case tcell.KeyDown:
				buttonPressed <- DOWN
			case tcell.KeyLeft:
				buttonPressed <- LEFT
			case tcell.KeyRight:
				buttonPressed <- RIGHT
			}
		}
	}
}

// Run is the function that starts the game
func Run(width int, heigth int, foodLimit int) {
	rand.Seed(time.Now().UnixNano())
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	screen.Clear()
	desk, snake, food := newGame(screen, width, heigth)
	food.limit = foodLimit
	keyEvents := make(chan button)

	go getEvents(screen, keyEvents)
	run := true
	for run {
		select {
		case bPressed := <-keyEvents:
			switch bPressed {
			case QUIT:
				run = false
			case RESTART:
				if !desk.running {
					desk, snake, food = restartGame(screen, desk, snake, food)
				}
			case UP, DOWN, LEFT, RIGHT:
				snake.direction = bPressed
			}
		case <-time.After(time.Millisecond * 50):
		}
		drawDesk(screen, desk)
		if desk.running {
			drawSnake(screen, desk, snake)
			addFood(food, snake, desk)
			drawFood(screen, desk, food)
			ateFood(food, snake, desk)
			checkCollison(snake, desk)
		}
		time.Sleep(time.Millisecond * time.Duration(100-5*desk.level))
	}
	screen.Fini()
}
