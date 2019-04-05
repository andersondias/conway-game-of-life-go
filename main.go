package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/andersondias/conway-game-of-life-go/conway"
)

func print(w conway.World) {
	for _, row := range w.Cells {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func main() {
	world := conway.World{}
	world.Setup(10, 10)
	world.Randomize()

	fmt.Println("\033[2J")
	reader := bufio.NewReader(os.Stdin)
	print(world)

	for {
		reader.ReadString('\n')
		fmt.Println("\033[2J")
		world.Tick()
		print(world)
	}
}
