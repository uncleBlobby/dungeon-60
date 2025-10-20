package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/world"
)

type PlayerSystem struct {
	World world.World
}

func (ps *PlayerSystem) Draw() {
	p := ps.World.GetPlayer()
	p.Draw()
}

func (ps *PlayerSystem) Update(dt float32) {
	p := ps.World.GetPlayer()
	movDir := p.RegisterMoveIntent()
	level := ps.World.GetLevel()

	move := rl.Vector2Scale(movDir, p.Speed*dt)

	p.Position.X += move.X
	p.UpdateCollider()
	p.ResolveWallCollision(level, movDir)

	p.Position.Y += move.Y
	p.UpdateCollider()
	p.ResolveWallCollision(level, movDir)

	p.Update(dt) // finally moves the player
}
