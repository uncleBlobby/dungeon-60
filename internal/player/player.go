package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/tile"
)

type Player struct {
	Position  rl.Vector2
	Direction rl.Vector2
	Speed     float32
	Collider  rl.Rectangle
}

func Create(x, y int) *Player {

	initMapPos := tile.Position{X: x, Y: y}
	initWorldPos := initMapPos.GetWorldPosition()

	p := &Player{
		Position:  initWorldPos,
		Direction: rl.Vector2{X: 0, Y: 0},
		Speed:     100,
	}

	p.Collider = rl.Rectangle{X: initWorldPos.X, Y: initWorldPos.Y, Width: tile.TILE_SIZE, Height: tile.TILE_SIZE}

	return p
}

func (p *Player) DebugDrawVelocity() {
	playerCenterX := int32(p.Position.X) + tile.TILE_SIZE/2
	playerCenterY := int32(p.Position.Y) + tile.TILE_SIZE/2
	playerVelocityX := playerCenterX + (int32(p.Direction.X) * 16)
	playerVelocityY := playerCenterY + (int32(p.Direction.Y) * 16)
	rl.DrawLine(playerCenterX, playerCenterY, playerVelocityX, playerVelocityY, rl.Black)
}

func (p *Player) DebugDrawTilePosition() {
	tilePos := p.GetTilePosition()
	rl.DrawRectangle(int32(tilePos.X)*tile.TILE_SIZE, int32(tilePos.Y)*tile.TILE_SIZE, tile.TILE_SIZE, tile.TILE_SIZE, rl.SkyBlue)
	//rl.DrawCircleV(tilePos.GetWorldPosition(), 5, rl.Red)
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), tile.TILE_SIZE, tile.TILE_SIZE, rl.Blue)
	p.DebugDrawTilePosition()
	p.DebugDrawVelocity()
}

func (p *Player) GetTilePosition() tile.Position {
	return tile.Position{
		X: int(math.Round(float64(p.Position.X+tile.TILE_SIZE/2))) / tile.TILE_SIZE,
		Y: int(math.Round(float64(p.Position.Y+tile.TILE_SIZE/2))) / tile.TILE_SIZE,
	}
}

func (p *Player) Update(dt float32) {

	p.Collider = rl.Rectangle{X: p.Position.X, Y: p.Position.Y, Width: tile.TILE_SIZE, Height: tile.TILE_SIZE}

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
