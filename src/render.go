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
			for j, cell := range universe[i] {
				raylib.DrawRectangle(
					int32(i)*w.CellSize,
					int32(j)*w.CellSize,
					w.CellSize,
					w.CellSize,
					func(alive bool) raylib.Color {
						if alive {
							return raylib.Black
						} else {
							return raylib.RayWhite
						}
					}(cell),
				)
			}
		}

		raylib.EndDrawing()
	}
}
