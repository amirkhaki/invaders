package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Space Invaders")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		var ground_height int32 = 50
		var score = int(rl.GetTime())
		var player_width int32 = 60

		// draw lives
		for i := range 3 {
			rl.DrawCircle(int32(30+50*i), 30, 20, rl.LightGray)
		}

		// draw score
		rl.DrawText(fmt.Sprintf("%.3d", score), int32(rl.GetScreenWidth())-80,
			30, 30, rl.DarkGray)

		// draw ground
		rl.DrawRectangleGradientH(0, int32(rl.GetScreenHeight())-ground_height,
			int32(rl.GetScreenWidth()), ground_height, rl.Brown, rl.LightGray)

		// draw player
		rl.DrawRectangle(int32(rl.GetScreenWidth()/2)-player_width/2,
			int32(rl.GetScreenHeight())-60, player_width, 20, rl.DarkPurple)

		//draw bases
		for i := range 3 {
			rl.DrawRectangle(int32(69+80*i), int32(rl.GetScreenHeight())/2+40,
				70, 70, rl.Beige)
		}

		rl.EndDrawing()
	}
}
