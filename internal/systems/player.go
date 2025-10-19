package systems

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/tile"
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

	// p.Direction = p.ResolveWallCollision(level, move)

	p.Position.Y += move.Y
	p.UpdateCollider()
	p.ResolveWallCollision(level, movDir)

	p.Update(dt) // finally moves the player
}

func (ps *PlayerSystem) UpdateOLD(dt float32) {
	p := ps.World.GetPlayer()

	movDir := p.RegisterMoveIntent()

	// check wall collision

	playerTile := p.GetTilePosition()

	playerNeighbours := playerTile.GetNeighbourPositionsCardinal()

	level := ps.World.GetLevel()

	for _, t := range level.Tilemap.Tiles {
		for _, n := range playerNeighbours {
			if t.Position.X == n.X && t.Position.Y == n.Y {
				worldPos := t.Position.GetWorldPosition()

				tt := level.Tilemap.Get(t.Position.X, t.Position.Y)

				transOrange := rl.ColorAlpha(rl.Orange, 0.5)
				transRed := rl.ColorAlpha(rl.Red, 0.5)

				if t.Walkable {
					rl.DrawRectangle(int32(worldPos.X), int32(worldPos.Y), tile.TILE_SIZE, tile.TILE_SIZE, transOrange)
				} else {
					rl.DrawRectangle(int32(worldPos.X), int32(worldPos.Y), tile.TILE_SIZE, tile.TILE_SIZE, transRed)
				}

				if rl.CheckCollisionRecs(p.Collider, tt.Collider) && !tt.Walkable {
					dir := rl.Vector2Subtract(tt.GetPositionAsVec2(), p.Position)
					if movDir.X != 0 || movDir.Y != 0 {
						fmt.Println("COLLISION DIRECTION: ", dir)
						fmt.Println("MOVE INTENT: ", movDir)
					}

					if dir.X > 0 && movDir.X > 0 {
						p.Position.X -= 1
						movDir.X = -0
					}

					if dir.Y > 0 && movDir.Y > 0 {
						p.Position.Y -= 1
						movDir.Y = -0
					}

					if dir.X < 0 && movDir.X < 0 {
						p.Position.X += 1
						movDir.X = 0
					}

					if dir.Y < 0 && movDir.Y < 0 {
						p.Position.Y += 1
						movDir.Y = 0
					}
				}
			}

		}
	}

	p.Direction = movDir

	p.Update(dt)
}
