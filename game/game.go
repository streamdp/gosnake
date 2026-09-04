package game

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
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
	PAUSE
)

type Game struct {
	screen    tcell.Screen
	desk      *desk
	snake     *snake
	food      *food
	keyEvents chan button
	running   bool
}

func NewGame(width int, height int, foodLimit int) *Game {
	g := newGame(newScreen(), width, height)
	g.food.limit = foodLimit

	go g.getEvents()

	return g
}

// Run is the function that starts the Game
func (g *Game) Run() {
	defer g.screen.Fini()

	for {
		select {
		case bPressed := <-g.keyEvents:
			switch bPressed {
			case QUIT:
				return
			case RESTART:
				if !g.isRunning() {
					g.restartGame()
				}
			case PAUSE:
				g.snake.paused = !g.snake.paused
				if g.snake.direction == 0 {
					g.snake.setDirection(getRandomDirection())
				}
			case UP, DOWN, LEFT, RIGHT:
				if g.snake.paused {
					g.snake.paused = false
				}

				if g.snake.direction == bPressed {
					continue
				}

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
		screen:    screen,
		desk:      d,
		snake:     newSnake(d.getRandPoint(), defaultSnakePalette),
		food:      newFood(),
		keyEvents: make(chan button),
		running:   true,
	}
}

func (g *Game) restartGame() *Game {
	g.screen.Clear()
	g.desk.level = 1
	g.desk.score = 0
	g.snake = newSnake(g.desk.getRandPoint(), defaultSnakePalette)
	g.food.position = make(map[coordinate]struct{}, g.food.limit)
	g.running = true

	return g
}

func (g *Game) getEvents() {
	for event := range g.screen.EventQ() {
		switch e := event.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventFocus:
			if !e.Focused && !g.snake.paused {
				g.keyEvents <- PAUSE
			}
		case *tcell.EventKey:
			switch e.Key() {
			case tcell.KeyEscape:
				g.keyEvents <- QUIT

				return
			case tcell.KeyEnter:
				g.keyEvents <- RESTART
			case tcell.KeyUp:
				g.keyEvents <- UP
			case tcell.KeyDown:
				g.keyEvents <- DOWN
			case tcell.KeyLeft:
				g.keyEvents <- LEFT
			case tcell.KeyRight:
				g.keyEvents <- RIGHT
			case tcell.KeyF2:
				g.keyEvents <- PAUSE
			default:
				continue
			}
		}
	}
}

func drawStr(screen tcell.Screen, x int, y int, style tcell.Style, str string, clr bool) {
	if clr {
		w, _ := screen.Size()
		for i := range w {
			if i >= x && i < x+len([]rune(str)) {
				continue
			}

			screen.SetContent(i, y, ' ', nil, style)
		}
	}

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
	for row := range g.desk.rect.height {
		for col := range g.desk.rect.width {
			if (row == 0 || row == g.desk.rect.height-1) || (col < 2 || col > g.desk.rect.width-3) {
				g.screen.SetContent(
					g.desk.rect.shiftX+col,
					g.desk.rect.shiftY+row,
					tcell.RuneCkBoard,
					nil,
					g.desk.palette.outer,
				)
			} else {
				g.screen.SetContent(
					g.desk.rect.shiftX+col,
					g.desk.rect.shiftY+row,
					' ',
					nil,
					g.desk.palette.inner,
				)
			}
		}
	}

	if !g.isRunning() {
		style := tcell.StyleDefault.Foreground(color.White).Background(color.DarkRed)
		text := "GAME OVER! YOU SCORE: " + strconv.Itoa(g.desk.score)
		drawStr(
			g.screen,
			g.desk.rect.shiftX+g.desk.rect.width/2-len([]rune(text))/2,
			g.desk.rect.height/2,
			style,
			text,
			false,
		)
		text = "PRESS ESC TO QUIT OR ENTER TO PLAY AGAIN"
		drawStr(
			g.screen,
			g.desk.rect.shiftX+g.desk.rect.width/2-len([]rune(text))/2,
			g.desk.rect.height/2+1,
			style.Reverse(true),
			text,
			false,
		)

		return
	}

	style := tcell.StyleDefault.Foreground(color.White).Background(color.Black)

	var msg string
	if g.snake.paused {
		msg = "     Game PAUSED (any arrows or F2 to unpause)"
	}

	text := fmt.Sprintf("Score: %d  Level: %d%s\r", g.desk.score, g.desk.level, msg)

	drawStr(g.screen, g.desk.rect.shiftX+1, 0, style, text, true)
}

func (g *Game) drawSnake() {
	g.snake.moveSnake()
	g.screen.SetContent(
		g.desk.rect.shiftX+g.snake.body[0].x,
		g.desk.rect.shiftY+g.snake.body[0].y,
		tcell.RuneCkBoard,
		nil,
		g.snake.palette.head,
	)

	for i := 1; i < g.snake.length; i++ {
		if i >= len(g.snake.body) {
			continue
		}

		g.screen.SetContent(
			g.desk.rect.shiftX+g.snake.body[i].x,
			g.desk.rect.shiftY+g.snake.body[i].y,
			tcell.RuneBoard,
			nil,
			g.snake.palette.body,
		)
	}
}

func (g *Game) drawFood() {
	style := tcell.StyleDefault.Background(color.DarkMagenta)
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

	if (g.snake.body[0].x == 1 || g.snake.body[0].x == g.desk.rect.width-2) ||
		(g.snake.body[0].y == 0 || g.snake.body[0].y == g.desk.rect.height-1) {
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

	screen.SetStyle(tcell.StyleDefault.Foreground(color.White).Background(color.Black))
	screen.EnableFocus()
	screen.Clear()

	return screen
}
