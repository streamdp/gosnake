package main

import (
	"flag"
	"fmt"

	"github.com/streamdp/gosnake/game"
)

func main()  {
	var showHelp bool
	var width int
	var heigth int
	var foodLimit int
	
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.IntVar(&width, "width", 70, "set width of the game desk")
	flag.IntVar(&heigth, "heigth", 20, "set heigth of the game desk")
	flag.IntVar(&foodLimit, "limit", 10, "set heigth of the game desk")
	flag.Parse()

	if showHelp {
		fmt.Println("gosnake is a version of the classic snake game written in golang with a library tcell.")
		fmt.Println("")
		flag.Usage()
		return
	}
	game.Run(width, heigth, foodLimit)
}


















