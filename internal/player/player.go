package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/tile"
)

type Player struct {
	Position  rl.Vector2
	Direction rl.Vector2
	Speed     float32
}

func Create() *Player {
	p := &Player{
		Position:  rl.Vector2{X: 0, Y: 0},
		Direction: rl.Vector2{X: 0, Y: 0},
		Speed:     100,
	}

	return p
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), tile.TILE_SIZE, tile.TILE_SIZE, rl.Blue)
}

func (p *Player) Update(dt float32) {

	p.Direction = rl.Vector2{X: 0, Y: 0}

	if rl.IsKeyDown(rl.KeyW) {
		p.Direction.Y = -1
	}

	if rl.IsKeyDown(rl.KeyS) {
		p.Direction.Y = +1
	}

	if rl.IsKeyDown(rl.KeyA) {
		p.Direction.X = -1
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.Direction.X = +1
	}

	p.Position.X += p.Direction.X * p.Speed * dt
	p.Position.Y += p.Direction.Y * p.Speed * dt
}
