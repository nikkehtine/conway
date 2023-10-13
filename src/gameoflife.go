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

	isAlive := func(x, y int) bool {
		var neighbors int
		for sampleX := -1; sampleX < 2; sampleX++ {
			for sampleY := -1; sampleY < 2; sampleY++ {
				if sampleX+x < 0 ||
					sampleX+x >= cols ||
					sampleY+y < 0 ||
					sampleY+y >= rows {
					continue
				}
				if sampleX == 0 && sampleY == 0 {
					continue
				}
				if world[sampleX+x][sampleY+y] {
					neighbors++
				}
			}
		}
		if world[x][y] && (2 <= neighbors && neighbors <= 3) {
			return true
		} else if !world[x][y] && neighbors == 3 {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			nextGen[i][j] = isAlive(i, j)
		}
	}

	return nextGen
}
