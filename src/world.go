package main

import "math/rand"

func make2DArray(cols int, rows int) Grid {
	arr := make(Grid, cols)
	for i := 0; i < cols; i++ {
		arr[i] = make([]bool, rows)
	}
	return arr
}

func randomizeGrid(grid Grid) *Grid {
	cols := len(grid)
	rows := len(grid[0])

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if rand.Intn(3) == 0 {
				grid[i][j] = true
			}
		}
	}
	return &grid
}
