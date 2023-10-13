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
	FPS:      12,
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

		for x := 0; x < WorldX; x++ {
			for y := 0; y < WorldY; y++ {
				raylib.DrawRectangle(
					int32(x)*w.CellSize,
					int32(y)*w.CellSize,
					w.CellSize,
					w.CellSize,
					func(alive bool) raylib.Color {
						if alive {
							return raylib.Black
						} else {
							return raylib.RayWhite
						}
					}(universe[x][y]),
				)
				raylib.DrawRectangleLines(
					int32(x)*w.CellSize,
					int32(y)*w.CellSize,
					w.CellSize+1,
					w.CellSize+1,
					raylib.Black,
				)
			}
		}

		universe = update(universe)

		raylib.EndDrawing()
	}
}
