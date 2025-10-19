package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/entity"
	"github.com/uncleBlobby/dungeon-60/internal/level"
	"github.com/uncleBlobby/dungeon-60/internal/player"
	"github.com/uncleBlobby/dungeon-60/internal/systems"
)

type Game struct {
	Level    *level.Level
	Player   *player.Player
	Entities []*entity.Entity

	LevelSystem  *systems.LevelSystem
	PlayerSystem *systems.PlayerSystem
	EntitySystem *systems.EntitySystem
	SpriteSystem *systems.SpriteSystem

	CollisionSystem *systems.CollisionSystem
}

func (g *Game) GetEntities() []*entity.Entity {
	return g.Entities
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
	g.Player = player.Create(46, 46)
	g.Entities = []*entity.Entity{}

	// g.Entities = append(g.Entities, entity.Create(10, 10, false, entity.ENTITY_BUSH))
	// g.Entities = append(g.Entities, entity.Create(20, 10, false, entity.ENTITY_TREE))

	g.LevelSystem = &systems.LevelSystem{World: g}
	g.PlayerSystem = &systems.PlayerSystem{World: g}
	g.EntitySystem = &systems.EntitySystem{World: g}
	g.SpriteSystem = systems.InitTextures(g)
	g.SpriteSystem.InitSprites()

	g.CollisionSystem = &systems.CollisionSystem{World: g}

	return g
}

func (g *Game) Draw() {

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	g.LevelSystem.Draw()
	g.PlayerSystem.Draw()
	g.EntitySystem.Draw()

	rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 5, 5, 24, rl.Black)
	//g.CollisionSystem.Draw()

	rl.EndDrawing()
}

func (g *Game) Update(dt float32) {
	g.LevelSystem.Update(dt)

	g.EntitySystem.Update(dt)

	g.PlayerSystem.Update(dt)
	// g.CollisionSystem.Update(dt)
}
