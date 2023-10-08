package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	raylib.InitWindow(w.Width, w.Height, w.Title)
	defer raylib.CloseWindow()
	raylib.SetTargetFPS(w.FPS)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.White)

		for i := range universe {
			for j, alive := range universe[i] {
				var cell_color raylib.Color
				if alive {
					cell_color = raylib.Black
				} else {
					cell_color = raylib.RayWhite
				}

				raylib.DrawRectangle(
					int32(i)*w.CellSize,
					int32(j)*w.CellSize,
					w.CellSize,
					w.CellSize,
					cell_color,
				)
			}
		}

		raylib.EndDrawing()
	}
}
