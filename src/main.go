package main

type Window struct {
	Width, Height int32
	Title         string
	FPS           int32
	CellSize      int32
}

var w = Window{
	Width:    800,
	Height:   450,
	Title:    "Game of Life",
	FPS:      30,
	CellSize: 8,
}

var universe Grid

func main() {
	cols := int(w.Width / w.CellSize)
	rows := int(w.Height / w.CellSize)

	universe = make2DArray(cols, rows)
	randomizeGrid(universe)
	draw()
}
