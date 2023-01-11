package game

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

type button int

// Game actions
const (
	RIGHT button = 1 + iota
	LEFT
	UP
	DOWN
	QUIT
	RESTART
)

type Game struct {
	screen  tcell.Screen
	desk    *desk
	snake   *snake
	food    *food
	running bool
}

func (g *Game) isRunning() bool {
	return g.running
}

func (g *Game) level() int {
	return g.desk.level
}

func newGame(screen tcell.Screen, width int, height int) *Game {
	screenSize, _ := screen.Size()
	d := newDesk(newRect(screenSize, width, height), defaultDeskPalette)
	return &Game{
		screen:  screen,
		desk:    d,
		snake:   newSnake(d.getRandPoint(), defaultSnakePalette),
		food:    newFood(),
		running: true,
	}
}

func (g *Game) restartGame() *Game {
	g.screen.Clear()
	g.desk.level = 0
	g.desk.score = 0
	g.snake.body = []coordinate{}
	g.snake.body = append(g.snake.body, g.desk.getRandPoint())
	g.snake.length = 3
	g.snake.direction = getRandomDirection()
	g.food.position = make(map[coordinate]struct{}, g.food.limit)
	g.running = true
	return g
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

func drawStr(screen tcell.Screen, x int, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		width := runewidth.RuneWidth(c)
		if width == 0 {
			comb = []rune{c}
			c = ' '
			width = 1
		}
		screen.SetContent(x, y, c, comb, style)
		x += width
	}
}

func (g *Game) drawDesk() {
	for row := 0; row < g.desk.rect.height; row++ {
		for col := 0; col < g.desk.rect.width; col++ {
			if (row == 0 || row == g.desk.rect.height-1) || (col < 2 || col > g.desk.rect.width-3) {
				g.screen.SetContent(g.desk.rect.shiftX+col, g.desk.rect.shiftY+row, tcell.RuneCkBoard, nil, g.desk.palette.outer)
			} else {
				g.screen.SetContent(g.desk.rect.shiftX+col, g.desk.rect.shiftY+row, rune(' '), nil, g.desk.palette.inner)
			}
		}
	}
	if !g.isRunning() {
		style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkRed)
		text := "GAME OVER! YOU SCORE: " + strconv.Itoa(g.desk.score)
		drawStr(g.screen, g.desk.rect.shiftX+g.desk.rect.width/2-len([]rune(text))/2, g.desk.rect.height/2, style, text)
		text = "PRESS ESC TO QUIT OR ENTER TO PLAY AGAIN"
		drawStr(g.screen, g.desk.rect.shiftX+g.desk.rect.width/2-len([]rune(text))/2, g.desk.rect.height/2+1, style.Reverse(true), text)
		return
	}
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	text := "Score: " + strconv.Itoa(g.desk.score) + "  Level: " + strconv.Itoa(g.desk.level)
	drawStr(g.screen, g.desk.rect.shiftX+1, 0, style, text)
}

func (g *Game) drawSnake() {
	g.snake.moveSnake()
	g.screen.SetContent(g.desk.rect.shiftX+g.snake.body[0].x, g.desk.rect.shiftY+g.snake.body[0].y, tcell.RuneCkBoard, nil, g.snake.palette.head)
	for i := 1; i < g.snake.length; i++ {
		g.screen.SetContent(g.desk.rect.shiftX+g.snake.body[i].x, g.desk.rect.shiftY+g.snake.body[i].y, tcell.RuneBoard, nil, g.snake.palette.body)
	}
}

func (g *Game) drawFood() {
	style := tcell.StyleDefault.Background(tcell.ColorDarkMagenta)
	for fp := range g.food.position {
		g.screen.SetContent(g.desk.rect.shiftX+fp.x, g.desk.rect.shiftY+fp.y, tcell.RuneCkBoard, nil, style)
	}
}

func (g *Game) checkCollisions() {
	for i := 1; i < len(g.snake.body); i++ {
		if g.snake.body[0] == g.snake.body[i] {
			g.running = false
		}
	}
	if (g.snake.body[0].x == 1 || g.snake.body[0].x == g.desk.rect.width-2) || (g.snake.body[0].y == 0 || g.snake.body[0].y == g.desk.rect.height-1) {
		g.running = false
	}
}

func newScreen() tcell.Screen {
	var (
		screen tcell.Screen
		err    error
	)
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
	return screen
}

func NewGame(width int, height int, foodLimit int) *Game {
	rand.Seed(time.Now().UnixNano())
	g := newGame(newScreen(), width, height)
	g.food.limit = foodLimit
	return g
}

// Run is the function that starts the Game
func (g *Game) Run() {
	defer g.screen.Fini()
	keyEvents := make(chan button)
	go getEvents(g.screen, keyEvents)
	for {
		select {
		case bPressed := <-keyEvents:
			switch bPressed {
			case QUIT:
				return
			case RESTART:
				if !g.isRunning() {
					g.restartGame()
				}
			case UP, DOWN, LEFT, RIGHT:
				g.snake.setDirection(bPressed)
			}
		case <-time.After(time.Millisecond * 50):
		}
		g.drawDesk()
		if g.isRunning() {
			g.drawSnake()
			g.food.add(g.snake, g.desk)
			g.drawFood()
			g.snake.ateFood(g.food, g.desk)
			g.checkCollisions()
		}
		g.screen.Show()
		time.Sleep(time.Millisecond * time.Duration(100-5*g.level()))
	}
}
