package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Grid struct {
	rackId     [300][300]int
	powerLevel [300][300]int
}

func (g *Grid) getGrid() (int, int) {
	highest := 0
	x, y := 0, 0
	for i := 0; i < 300-2; i++ {
		for j := 0; j < 300-2; j++ {
			sum := 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					sum += g.powerLevel[i+k][j+l]
				}
			}
			if sum > highest {
				highest = sum
				x = i
				y = j
			}
		}
	}
	return PuzzleCoordinates(x, y)
}

func (g *Grid) getGridPart2() (int, int, int) {
	highest := 0
	x, y, size := 0, 0, 0

	for s := 1; s <= 300; s++ {
		for i := 0; i < 300-s+1; i++ {
			for j := 0; j < 300-s+1; j++ {
				sum := 0
				for k := 0; k < s; k++ {
					for l := 0; l < s; l++ {
						sum += g.powerLevel[i+k][j+l]
					}
				}
				if sum > highest {
					highest = sum
					x = i
					y = j
					size = s
				}
			}
		}
	}
	xPuzzle, yPuzzle := PuzzleCoordinates(x, y)
	return xPuzzle, yPuzzle, size
}

func (g *Grid) setRackId() {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			x, _ := PuzzleCoordinates(i, j)
			g.rackId[i][j] = x + 10
		}
	}
}

func (g *Grid) setBasicPowerLevel() {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			_, y := PuzzleCoordinates(i, j)
			g.powerLevel[i][j] = y * g.rackId[i][j]
		}
	}
}

func (g *Grid) increasePowerLevel(serial int) {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			g.powerLevel[i][j] += serial
		}
	}
}

func (g *Grid) multiplyPowerLevel() {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			g.powerLevel[i][j] *= g.rackId[i][j]
		}
	}
}

func (g *Grid) reducePowerLevel() {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			// inp := g.powerLevel[i][j]
			s := strconv.Itoa(g.powerLevel[i][j])
			digits := []rune(s)
			if len(digits) >= 3 {
				r := digits[len(digits)-3]
				if r >= '0' && r <= '9' {
					g.powerLevel[i][j] = int(r - '0')
				} else {
					g.powerLevel[i][j] = 0
				}
			} else {
				g.powerLevel[i][j] = 0
			}
			// out := g.powerLevel[i][j]
			// fmt.Println(inp, out)
		}
	}
}

func (g *Grid) substractPowerLevel() {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			g.powerLevel[i][j] -= 5
		}
	}
}

func PuzzleCoordinates(x, y int) (int, int) {
	return x + 1, y + 1
}

func Part1(serial int) (int, int) {
	grid := Grid{}
	grid.setRackId()
	grid.setBasicPowerLevel()
	grid.increasePowerLevel(serial)
	grid.multiplyPowerLevel()
	grid.reducePowerLevel()
	grid.substractPowerLevel()
	return grid.getGrid()
}

func Part2(serial int) (int, int, int) {
	grid := Grid{}
	grid.setRackId()
	grid.setBasicPowerLevel()
	grid.increasePowerLevel(serial)
	grid.multiplyPowerLevel()
	grid.reducePowerLevel()
	grid.substractPowerLevel()
	return grid.getGridPart2()
}

func main() {
	serial, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	x, y := Part1(serial)
	outputPart1 := fmt.Sprintf("%d,%d", x, y)
	fmt.Println(outputPart1)

	x, y, size := Part2(serial)
	outputPart2 := fmt.Sprintf("%d,%d,%d", x, y, size)
	fmt.Println(outputPart2)
}
