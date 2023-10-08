package main

type Grid [][]bool
type Window struct {
	Width, Height int32
	Title         string
	FPS           int32
	CellSize      int32
}

var currentGen, nextGen Grid
var w = Window{
	Width:    800,
	Height:   450,
	Title:    "Game of Life",
	FPS:      12,
	CellSize: 8,
}

func main() {
	cols := int(w.Width / w.CellSize)
	rows := int(w.Height / w.CellSize)

	currentGen = make2DArray(cols, rows)
	randomizeGrid(currentGen)
	nextGen = make2DArray(cols, rows)

	draw()
}
