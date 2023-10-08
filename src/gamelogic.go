package main

import (
	"log"
	"math/rand"
)

var current_cell bool
var next_cell bool
var neighbors int // 0...8

type Grid [][]bool

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

func update_cell() {
	if neighbors < 0 || neighbors > 8 {
		log.Fatal("Neighbors out of bounds!")
	}
	if current_cell {
		if neighbors == 3 {
			next_cell = true
		} else {
			next_cell = false
		}
	} else {
		if 2 <= neighbors && neighbors <= 3 {
			next_cell = true
		} else {
			next_cell = false
		}
	}
}
