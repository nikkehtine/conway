package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Grid [][]bool

const (
	APP_WIDTH, APP_HEIGHT int32  = 800, 450
	CELL_SIZE             int32  = 8
	FPS                   int32  = 18 // The higher the value, the faster the game
	APP_TITLE             string = "Game of Life"
)

var universe Grid

func main() {
	WorldX := int(APP_WIDTH / CELL_SIZE)
	WorldY := int(APP_HEIGHT / CELL_SIZE)

	universe = make2DArray(WorldX, WorldY)
	universe = randomizeWorld(universe) // Randomize by default (for now)

	// It's raylibbing time

	raylib.InitWindow(APP_WIDTH, APP_HEIGHT, APP_TITLE)
	defer raylib.CloseWindow()
	raylib.SetTargetFPS(FPS)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.White)

		for col := 0; col < WorldX; col++ {
			for row := 0; row < WorldY; row++ {
				size := CELL_SIZE
				x, y := int32(col)*size, int32(row)*size
				color, lnColor := raylib.RayWhite, raylib.Black
				if universe[col][row] {
					color = raylib.Black
				}
				raylib.
					DrawRectangle(x, y, size, size, color)
				raylib.
					DrawRectangleLines(x, y, size+1, size+1, lnColor)
			}
		}

		universe = update(universe)

		raylib.EndDrawing()
	}
}
