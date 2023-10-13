package main

import (
	"math/rand"
)

func make2DArray(cols int, rows int) Grid {
	world := make(Grid, cols)
	for i := 0; i < cols; i++ {
		world[i] = make([]bool, rows)
	}
	return world
}

func randomizeWorld(world Grid) Grid {
	cols, rows := len(world), len(world[0])
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if rand.Intn(3) == 0 {
				world[i][j] = true
			}
		}
	}
	return world
}

func update(world Grid) Grid {
	cols, rows := len(world), len(world[0])
	nextGen := make2DArray(cols, rows)

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			var neighbors int
			for sampleX := -1; sampleX < 2; sampleX++ {
				for sampleY := -1; sampleY < 2; sampleY++ {
					if sampleX+i < 0 ||
						sampleX+i >= cols ||
						sampleY+j < 0 ||
						sampleY+j >= rows {
						continue
					}
					if sampleX == 0 && sampleY == 0 {
						continue
					}
					if world[sampleX+i][sampleY+j] {
						neighbors++
					}
				}
			}

			if neighbors == 3 ||
				(world[i][j] && (2 <= neighbors && neighbors <= 3)) {
				nextGen[i][j] = true
			} else {
				nextGen[i][j] = false
			}
		}
	}

	return nextGen
}
