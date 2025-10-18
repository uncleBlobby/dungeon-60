package tile

import rl "github.com/gen2brain/raylib-go/raylib"

const TILE_SIZE = 16

type TileType int

const (
	TILETYPE_STONE TileType = iota
	TILETYPE_GRASS
)

type Position struct {
	X int
	Y int
}

type Tile struct {
	Position Position
	Size     int
	Type     TileType
}

func Create(posX, posY int, size int, t TileType) *Tile {
	tile := &Tile{
		Position: Position{
			X: posX,
			Y: posY,
		},
		Size: size,
		Type: t,
	}

	return tile
}

func (t *Tile) GetPosition() Position {
	return t.Position
}

func (t *Tile) GetPositionAsVec2() rl.Vector2 {
	return rl.Vector2{X: float32(t.Position.X) * TILE_SIZE, Y: float32(t.Position.Y) * TILE_SIZE}
}

func (t *Tile) GetWorldPosition() rl.Vector2 {
	return t.GetPositionAsVec2()
}

func (t *Tile) GetDrawColor() rl.Color {
	switch t.Type {
	case TILETYPE_STONE:
		return rl.LightGray
	case TILETYPE_GRASS:
		return rl.Green
	default:
		return rl.Red
	}
}

func (t *Tile) Draw() {
	drawCoords := t.GetWorldPosition()
	rl.DrawRectangle(int32(drawCoords.X), int32(drawCoords.Y), int32(t.Size), int32(t.Size), t.GetDrawColor())
}
