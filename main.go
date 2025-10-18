package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/game"
)

func main() {
	fmt.Println("Hello, dungeon-60!")

	rl.InitWindow(800, 600, "dungeon-60 v0.1")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	g := game.Create()

	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()

		g.Update(dt)

		g.Draw()

	}
}
