package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/level"
	"github.com/uncleBlobby/dungeon-60/internal/player"
	"github.com/uncleBlobby/dungeon-60/internal/systems"
)

type Game struct {
	Level  *level.Level
	Player *player.Player

	LevelSystem  *systems.LevelSystem
	PlayerSystem *systems.PlayerSystem
}

func (g *Game) GetLevel() *level.Level {
	return g.Level
}

func (g *Game) GetPlayer() *player.Player {
	return g.Player
}

func Create() *Game {
	g := &Game{}

	g.Level = level.Create(128, 128)
	g.Player = player.Create()

	g.LevelSystem = &systems.LevelSystem{World: g}
	g.PlayerSystem = &systems.PlayerSystem{World: g}

	return g
}

func (g *Game) Draw() {

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	g.LevelSystem.Draw()
	g.PlayerSystem.Draw()

	rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 5, 5, 24, rl.Black)

	rl.EndDrawing()
}

func (g *Game) Update(dt float32) {
	g.LevelSystem.Update(dt)
	g.PlayerSystem.Update(dt)
}
