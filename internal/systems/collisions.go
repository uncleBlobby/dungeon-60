package systems

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/tile"
	"github.com/uncleBlobby/dungeon-60/internal/world"
)

type CollisionSystem struct {
	World world.World
}

func (cs *CollisionSystem) Update(dt float32) { // UNUSED
	p := cs.World.GetPlayer()

	playerTile := p.GetTilePosition()
	fmt.Println(playerTile)

	playerNeighbours := playerTile.GetNeighbourPositionsCardinal()
	fmt.Println(playerNeighbours)

	level := cs.World.GetLevel()

	for _, t := range level.Tilemap.Tiles {
		for _, n := range playerNeighbours {
			if t.Position.X == n.X && t.Position.Y == n.Y {
				fmt.Println("NEIGHBOUR FOUND")
				worldPos := t.Position.GetWorldPosition()
				fmt.Println(p.Position)
				fmt.Println(worldPos)
				rl.DrawRectangle(int32(worldPos.X), int32(worldPos.Y), tile.TILE_SIZE, tile.TILE_SIZE, rl.Orange)
			}
		}
		// if t.Position
	}

	// get tiles surrounding player

}

func (cs *CollisionSystem) Draw() { // UNUSED (DEBUGGING)
	p := cs.World.GetPlayer()

	playerTile := p.GetTilePosition()
	fmt.Println(playerTile)

	playerNeighbours := playerTile.GetNeighbourPositionsAll()
	fmt.Println(playerNeighbours)

	level := cs.World.GetLevel()

	for _, t := range level.Tilemap.Tiles {
		for _, n := range playerNeighbours {
			if t.Position.X == n.X && t.Position.Y == n.Y {
				// fmt.Println("NEIGHBOUR FOUND")
				worldPos := t.Position.GetWorldPosition()
				// fmt.Println(p.Position)
				// fmt.Println(worldPos)

				tt := level.Tilemap.Get(t.Position.X, t.Position.Y)

				transOrange := rl.ColorAlpha(rl.Orange, 0.5)
				transRed := rl.ColorAlpha(rl.Red, 0.5)

				if t.Walkable {
					rl.DrawRectangle(int32(worldPos.X), int32(worldPos.Y), tile.TILE_SIZE, tile.TILE_SIZE, transOrange)
				} else {
					rl.DrawRectangle(int32(worldPos.X), int32(worldPos.Y), tile.TILE_SIZE, tile.TILE_SIZE, transRed)
				}

				// check player colliding neighbour

				if rl.CheckCollisionRecs(p.Collider, tt.Collider) && !tt.Walkable {
					rl.DrawRectangle(tt.Collider.ToInt32().X, tt.Collider.ToInt32().Y, tt.Collider.ToInt32().Width, tt.Collider.ToInt32().Height, rl.Red)

					// if it is in the X direction,

					dir := rl.Vector2Subtract(p.Position, tt.GetPositionAsVec2())

					if dir.X < 0 || dir.X > 0 {
						p.Direction.X = 0
					}
				}

			}

		}
		// if t.Position
	}

	// get tiles surrounding player
}
