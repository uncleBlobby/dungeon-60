package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/level"
	"github.com/uncleBlobby/dungeon-60/internal/tile"
)

type Player struct {
	Position  rl.Vector2
	Direction rl.Vector2
	Facing    rl.Vector2
	Speed     float32
	Collider  rl.Rectangle
}

func Create(x, y int) *Player {

	initMapPos := tile.Position{X: x, Y: y}
	initWorldPos := initMapPos.GetWorldPosition()

	p := &Player{
		Position:  initWorldPos,
		Direction: rl.Vector2{X: 0, Y: 0},
		Facing:    rl.Vector2{X: 0, Y: 0},
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

func (p *Player) DrawEquipmentAttachment() {
	mousePos := rl.GetMousePosition()
	radius := float32(32)
	playerCenter := rl.Vector2{X: p.Position.X + tile.TILE_SIZE/2, Y: p.Position.Y + tile.TILE_SIZE/2}

	// Get direction vector from player to mouse
	mousedir := rl.Vector2Subtract(mousePos, playerCenter)

	// Normalize that direction
	mousedir = rl.Vector2Normalize(mousedir)

	// Scale that direction by the radius (how far out is the equipment attachment point?)
	offset := rl.Vector2Scale(mousedir, radius)

	// Add that offset vector to the player (center) position
	circlePoint := rl.Vector2Add(playerCenter, offset)
	rl.DrawCircleV(circlePoint, 10, rl.Red)

	// Draw radius for debug
	rl.DrawCircleLinesV(playerCenter, radius, rl.Blue)
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), tile.TILE_SIZE, tile.TILE_SIZE, rl.Blue)
	// p.DebugDrawVelocity()
	p.DrawEquipmentAttachment()
}

func (p *Player) GetTilePosition() tile.Position {
	return tile.Position{
		X: int(math.Round(float64(p.Position.X+tile.TILE_SIZE/2))) / tile.TILE_SIZE,
		Y: int(math.Round(float64(p.Position.Y+tile.TILE_SIZE/2))) / tile.TILE_SIZE,
	}
}

func (p *Player) UpdateCollider() {
	p.Collider = rl.Rectangle{X: p.Position.X, Y: p.Position.Y, Width: tile.TILE_SIZE, Height: tile.TILE_SIZE}
}

func (p *Player) Update(dt float32) {

	// TODO
	// update facing direction to follow mouse pointer

	p.Position.X += p.Direction.X * p.Speed * dt
	p.Position.Y += p.Direction.Y * p.Speed * dt
	p.UpdateCollider()
}

func (p *Player) RegisterMoveIntent() rl.Vector2 {

	movIntent := rl.Vector2{X: 0, Y: 0}

	if rl.IsKeyDown(rl.KeyW) {
		movIntent.Y = -1
	}

	if rl.IsKeyDown(rl.KeyS) {
		movIntent.Y = +1
	}

	if rl.IsKeyDown(rl.KeyA) {
		movIntent.X = -1
	}

	if rl.IsKeyDown(rl.KeyD) {
		movIntent.X = +1
	}

	return movIntent
}

func (p *Player) CollidesWithWall(level *level.Level) bool {

	// TODO
	// Optimize routine so it only loops through tiles next to player (neighbours)
	// rather than every tile in the level -- for now it works okay
	for _, t := range level.Tilemap.Tiles {
		if rl.CheckCollisionRecs(p.Collider, t.Collider) && !t.Walkable {
			return true
		}
	}
	return false
}

func (p *Player) ResolveWallCollision(level *level.Level, movDir rl.Vector2) rl.Vector2 {
	for _, t := range level.Tilemap.Tiles {
		if rl.CheckCollisionRecs(p.Collider, t.Collider) && !t.Walkable {
			overlapX := getOverlapX(p.Collider, t.Collider)
			overlapY := getOverlapY(p.Collider, t.Collider)

			if overlapX > 0 && overlapY > 0 {

				// Correct along the smallest overlap axis to avoid diagonal push
				if overlapX < overlapY {
					if p.Collider.X < t.Collider.X {
						// Player is to the LEFT of tile
						p.Position.X -= overlapX
					} else {
						// Player is to the RIGHT of tile
						p.Position.X += overlapX
					}
				} else {
					if p.Collider.Y < t.Collider.Y {
						// Player is ABOVE tile (remember: Y increases downward)
						p.Position.Y -= overlapY
					} else {
						// Player is BELOW tile
						p.Position.Y += overlapY
					}
				}
				p.UpdateCollider()
			}
		}
	}
	return movDir
}

func getOverlapX(a, b rl.Rectangle) float32 {
	left := float32(math.Max(float64(a.X), float64(b.X)))
	right := float32(math.Min(float64(a.X+a.Width), float64(b.X+b.Width)))
	return right - left
}

func getOverlapY(a, b rl.Rectangle) float32 {
	top := float32(math.Max(float64(a.Y), float64(b.Y)))
	bottom := float32(math.Min(float64(a.Y+a.Height), float64(b.Y+b.Height)))
	return bottom - top
}
