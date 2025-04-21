package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "math"

const MAX_FPS = 12
const WINDOW_X = 750
const WINDOW_Y = 759
const CELL_SIZE = 25
const WINDOW_TITLE = "raylib game of life"

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagBorderlessWindowedMode)
	rl.SetTargetFPS(MAX_FPS)
	rl.InitWindow(WINDOW_X, WINDOW_Y, WINDOW_TITLE)
	defer rl.CloseWindow()

	simulation := InitSimulation(WINDOW_X, WINDOW_Y, CELL_SIZE)

	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mousePosition := rl.GetMousePosition()
			row := int32(mousePosition.Y / CELL_SIZE)
			column := int32(mousePosition.X / CELL_SIZE)
			simulation.ToggleCell(row, column)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			if simulation.IsPaused() {
				simulation.Resume()
			} else {
				simulation.Pause()
			}
		}

		if rl.IsKeyPressed(rl.KeyC) {
			simulation.Clear()
		}

		if rl.IsKeyPressed(rl.KeyR) {
			simulation.FillRandom()
		}

		simulation.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		simulation.Draw()

		rl.EndDrawing()
	}
}

func HandleCameraZoom(camera rl.Camera2D) {

	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
		camera.Offset = rl.GetMousePosition()
		camera.Target = mouseWorldPos
		scale := 0.05 * wheel
		calc := math.Exp(math.Log(float64(camera.Zoom) + float64(scale)))
		val := rl.Clamp(float32(calc), 0.125, 0.64)
		camera.Zoom = val
	}

}
