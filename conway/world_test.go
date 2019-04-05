package conway_test

import (
	"testing"

	"github.com/andersondias/conway-game-of-life-go/conway"
)

func evaluateCell(t *testing.T, w conway.World, x int, y int, expectedStatus int) {
	if w.Cells[x][y] != expectedStatus {
		t.Errorf("cell at %d x %d should be %d", x, y, expectedStatus)
	}
}

func TestSetup(t *testing.T) {
	world := conway.World{}
	world.Setup(2, 3)

	height := len(world.Cells)

	if height != 2 {
		t.Errorf("height = %d; want 2", height)
	}

	width := len(world.Cells[0])
	if width != 3 {
		t.Errorf("width = %d; want 3", width)
	}
}

func TestTickFirstRule(t *testing.T) {
	world := conway.World{}
	world.Setup(3, 3)
	world.Cells[0][0] = 1
	world.Cells[1][1] = 0
	world.Cells[2][2] = 0

	world.Tick()

	evaluateCell(t, world, 0, 0, conway.Dead)
	evaluateCell(t, world, 1, 1, conway.Dead)
	evaluateCell(t, world, 2, 2, conway.Dead)
}

func TestTickSecondRule(t *testing.T) {
	world := conway.World{}
	world.Setup(3, 3)
	world.Cells[0][0] = 1
	world.Cells[1][0] = 1
	world.Cells[1][2] = 1
	world.Cells[2][0] = 1
	world.Cells[2][1] = 1

	world.Tick()

	evaluateCell(t, world, 0, 0, conway.Dead)
	evaluateCell(t, world, 1, 0, conway.Alive)
	evaluateCell(t, world, 1, 2, conway.Dead)
	evaluateCell(t, world, 2, 0, conway.Alive)
	evaluateCell(t, world, 2, 1, conway.Alive)
}

func TestTickThirdRule(t *testing.T) {
	world := conway.World{}
	world.Setup(3, 3)
	world.Cells[0][0] = 1
	world.Cells[1][0] = 1
	world.Cells[1][1] = 1
	world.Cells[1][2] = 1
	world.Cells[2][0] = 1
	world.Cells[2][1] = 1
	world.Cells[2][2] = 1

	world.Tick()

	evaluateCell(t, world, 0, 0, conway.Alive)
	evaluateCell(t, world, 1, 0, conway.Dead)
	evaluateCell(t, world, 1, 1, conway.Dead)
	evaluateCell(t, world, 1, 2, conway.Alive)
	evaluateCell(t, world, 2, 0, conway.Alive)
	evaluateCell(t, world, 2, 1, conway.Dead)
	evaluateCell(t, world, 2, 2, conway.Alive)
}

func TestTickForthRule(t *testing.T) {
	world := conway.World{}
	world.Setup(3, 3)
	world.Cells[0][0] = 1
	world.Cells[0][2] = 1
	world.Cells[1][1] = 1
	world.Cells[2][1] = 1

	world.Tick()

	evaluateCell(t, world, 0, 1, conway.Alive)
	evaluateCell(t, world, 1, 0, conway.Alive)
	evaluateCell(t, world, 1, 2, conway.Alive)
	evaluateCell(t, world, 2, 0, conway.Dead)
	evaluateCell(t, world, 2, 2, conway.Dead)
}
