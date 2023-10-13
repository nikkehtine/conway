package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Grid [][]bool
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
	FPS:      18,
	CellSize: 8,
}

var universe Grid

func main() {
	WorldX := int(w.Width / w.CellSize)
	WorldY := int(w.Height / w.CellSize)

	universe = make2DArray(WorldX, WorldY)
	universe = randomizeWorld(universe) // Randomize by default (for now)

	// It's raylibbing time

	raylib.InitWindow(w.Width, w.Height, w.Title)
	defer raylib.CloseWindow()
	raylib.SetTargetFPS(w.FPS)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.White)

		for col := 0; col < WorldX; col++ {
			for row := 0; row < WorldY; row++ {
				size := w.CellSize
				x, y := int32(col)*size, int32(row)*size
				color, lnColor := raylib.RayWhite, raylib.Black
				if universe[col][row] {
					color = raylib.Black
				}
				raylib.DrawRectangle(x, y, size, size, color)
				raylib.DrawRectangleLines(x, y, size+1, size+1, lnColor)
			}
		}

		universe = update(universe)

		raylib.EndDrawing()
	}
}
