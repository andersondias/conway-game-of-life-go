package conway

import (
	"math/rand"
	"time"
)

// Alive a
const Alive = 1

// Dead a
const Dead = 0

// World alguma coisa
type World struct {
	Cells [][]int
}

// Setup alguma coisa
func (w *World) Setup(height, width int) {
	w.Cells = make([][]int, height)
	for row := range w.Cells {
		w.Cells[row] = make([]int, width)
	}
}

// Randomize a
func (w *World) Randomize() {
	rand.Seed(time.Now().UnixNano())

	for _, row := range w.Cells {
		for c := range row {
			row[c] = rand.Intn(2)
		}
	}
}

// Tick alguma coisa
func (w *World) Tick() {
	var cellsToKill, cellsToRessurrect [][]int
	for r, row := range w.Cells {
		for c, cell := range row {
			aliveCount := w.aliveSurroundingCells(r, c)
			if cell == Alive {
				if aliveCount != 2 && aliveCount != 3 {
					cellsToKill = append(cellsToKill, []int{r, c})
				}
			} else if aliveCount == 3 {
				cellsToRessurrect = append(cellsToRessurrect, []int{r, c})
			}
		}
	}
	for _, cellToKill := range cellsToKill {
		w.Cells[cellToKill[0]][cellToKill[1]] = 0
	}
	for _, cellToRessurrect := range cellsToRessurrect {
		w.Cells[cellToRessurrect[0]][cellToRessurrect[1]] = 1
	}
}

func (w World) surroundingCells(x, y int) []int {
	var surroundingCells []int
	var top int
	var down int
	var left int
	var right int

	top = x - 1
	down = x + 1
	left = y - 1
	if left == -1 {
		left = y
	}

	right = y + 1
	if right == len(w.Cells[x]) {
		right = y
	}

	if top >= 0 && top != x {
		surroundingCells = append(surroundingCells, w.Cells[top][left:right+1]...)
	}

	if left != y {
		surroundingCells = append(surroundingCells, w.Cells[x][left])
	}
	if right != y {
		surroundingCells = append(surroundingCells, w.Cells[x][right])
	}

	if down < len(w.Cells) {
		surroundingCells = append(surroundingCells, w.Cells[down][left:right+1]...)
	}

	return surroundingCells
}

func (w World) aliveSurroundingCells(x, y int) int {
	var aliveCount int
	for _, cell := range w.surroundingCells(x, y) {
		if cell == Alive {
			aliveCount++
		}
	}
	return aliveCount
}
