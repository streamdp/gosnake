package main

import (
	"flag"
	"fmt"

	"github.com/streamdp/gosnake/game"
)

func main() {
	var (
		showHelp         bool
		width            int
		height           int
		foodLimit        int
		adjustableScreen bool
	)

	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.IntVar(&width, "width", 70, "set width of the game desk")
	flag.IntVar(&height, "height", 20, "set height of the game desk")
	flag.IntVar(&foodLimit, "limit", 10, "set food limit")
	flag.BoolVar(&adjustableScreen, "adjustable-screen", true,
		"screen size can be automatically adjusted to the terminal window size",
	)

	flag.Parse()

	if showHelp {
		fmt.Println("gosnake is a version of the classic snake game written in golang with a library tcell.")
		fmt.Println("")
		flag.Usage()

		return
	}

	game.NewGame(width, height, foodLimit, adjustableScreen).Run()
}
