package game

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

type button int

const (
	QUIT button = 1 + iota
	RESTART
	B_RIGHT
	B_LEFT
	B_UP
	B_DOWN
)

func newGame(screen tcell.Screen, width int, heigth int) (desk *desk, snake *snake, food *food) {
	desk = newDesk(newRect(screen, width, heigth), &deskPallete{
		inner: tcell.StyleDefault.Background(tcell.ColorBisque),
		outer: tcell.StyleDefault.Background(tcell.ColorPaleVioletRed),
	})
	x := rand.Intn(desk.rect.width-3) + 2
	y := rand.Intn(desk.rect.heigth-2) + 1
	snake = newSnake(x, y, &snakePallete{
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
	snake.body = append(snake.body, coord{
		x: 2 + rand.Intn(desk.rect.width-3),
		y: 1 + rand.Intn(desk.rect.heigth-2),
	})
	snake.length = 3
	snake.direction = 1 + direction(rand.Intn(3))
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
				buttonPressed <- B_UP
			case tcell.KeyDown:
				buttonPressed <- B_DOWN
			case tcell.KeyLeft:
				buttonPressed <- B_LEFT
			case tcell.KeyRight:
				buttonPressed <- B_RIGHT
			}
		}
	}
}

func Run(width int, heigth int, foodLimit int) {
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
			case B_UP:
				snake.direction = UP
			case B_DOWN:
				snake.direction = DOWN
			case B_LEFT:
				snake.direction = LEFT
			case B_RIGHT:
				snake.direction = RIGHT
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
