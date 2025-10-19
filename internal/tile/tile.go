package tile

import rl "github.com/gen2brain/raylib-go/raylib"

const TILE_SIZE = 16

type TileType int

const (
	TILETYPE_STONE TileType = iota
	TILETYPE_GRASS
	TILETYPE_DIRT
)

type Position struct {
	X int
	Y int
}

func (p *Position) GetNeighbourPositionsAll() []Position {
	directions := []Position{
		//left
		{X: -1, Y: 0},
		//right
		{X: 1, Y: 0},
		//down
		{X: 0, Y: 1},
		//up
		{X: 0, Y: -1},
		//left-up
		{X: -1, Y: -1},
		//right-up
		{X: 1, Y: -1},
		//left-down
		{X: -1, Y: 1},
		//right-down
		{X: 1, Y: 1},
	}

	nbs := []Position{}

	for _, dir := range directions {
		nbs = append(nbs, Position{
			X: p.X + dir.X,
			Y: p.Y + dir.Y,
		})
	}

	return nbs
}

func (p *Position) GetNeighbourPositionsCardinal() []Position {
	directions := []Position{
		//left
		{X: -1, Y: 0},
		//right
		{X: 1, Y: 0},
		//down
		{X: 0, Y: 1},
		//up
		{X: 0, Y: -1},
	}

	nbs := []Position{}

	for _, dir := range directions {
		nbs = append(nbs, Position{
			X: p.X + dir.X,
			Y: p.Y + dir.Y,
		})
	}

	return nbs
}

type Tile struct {
	Position Position
	Size     int
	Type     TileType
	Collider rl.Rectangle
	Walkable bool
}

func (t *Tile) SetDefaultWalkable() {
	switch t.Type {
	case TILETYPE_STONE:
		t.Walkable = false
	default:
		t.Walkable = true
	}
}

func Create(posX, posY int, size int, t TileType) *Tile {
	pos := Position{
		X: posX,
		Y: posY,
	}
	worldPos := pos.GetWorldPosition()
	tile := &Tile{
		Position: Position{
			X: posX,
			Y: posY,
		},
		Size:     size,
		Type:     t,
		Collider: rl.Rectangle{X: worldPos.X, Y: worldPos.Y, Width: TILE_SIZE, Height: TILE_SIZE},
	}

	tile.SetDefaultWalkable()
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

func (p *Position) GetWorldPosition() rl.Vector2 {
	return rl.Vector2{X: float32(p.X) * TILE_SIZE, Y: float32(p.Y) * TILE_SIZE}
}

func (t *Tile) GetDrawColor() rl.Color {
	switch t.Type {
	case TILETYPE_STONE:
		return rl.Gray
	case TILETYPE_GRASS:
		return rl.Green
	case TILETYPE_DIRT:
		return rl.Brown
	default:
		return rl.Red
	}
}

func (t *Tile) Draw() {
	drawCoords := t.GetWorldPosition()
	rl.DrawRectangle(int32(drawCoords.X), int32(drawCoords.Y), int32(t.Size), int32(t.Size), t.GetDrawColor())
}
