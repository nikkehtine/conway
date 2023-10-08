package main

import (
	"log"
)

func updateWorld() {
	for i := range currentGen {
		for j := range currentGen[i] {
			nextGen[i][j] = updateCell(
				currentGen[i][j],
				countNeighbors(i, j),
			)
		}
	}
	currentGen = nextGen
}

func updateCell(currentCell bool, neighbors int) bool {
	if neighbors < 0 || neighbors > 8 {
		log.Fatal("Neighbors out of bounds!")
	}

	if currentCell {
		return 2 <= neighbors && neighbors <= 3
	} else {
		return neighbors == 3
	}
}

func countNeighbors(x, y int) int {
	var neighbors int
	for i := -1; i < 2; i++ {
		if x+i < 0 || x+i >= len(currentGen) {
			continue
		}
		for j := -1; j < 2; j++ {
			if y+j < 0 || y+j >= len(currentGen[0]) {
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			if currentGen[x+i][y+j] {
				neighbors++
			}
		}
	}
	return neighbors
}
