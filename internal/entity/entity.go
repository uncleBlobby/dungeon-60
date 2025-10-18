package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/tile"
)

type Entity struct {
	Position tile.Position
	Sprite   rl.Texture2D
	Solid    bool
	Type     EntityType
}

type EntityType string

const (
	ENTITY_TREE EntityType = "tree"
	ENTITY_BUSH EntityType = "bush"
)

func Create(x, y int, s bool, t EntityType) *Entity {

	e := &Entity{
		Position: tile.Position{
			X: x,
			Y: y,
		},
		Solid: false,
		Type:  t,
	}

	if e.Type == ENTITY_BUSH {
		e.Sprite = rl.LoadTexture("assets/sprites/bush1.png")
	}

	return e
}

func (e *Entity) Draw() {

	gridPos := e.Position.GetWorldPosition()

	if e.Type == ENTITY_BUSH {
		rl.DrawTexture(e.Sprite, int32(gridPos.X), int32(gridPos.Y), rl.White)
	} else {
		rl.DrawRectangle(int32(gridPos.X), int32(gridPos.Y), tile.TILE_SIZE, tile.TILE_SIZE, rl.Red)
	}
}

func (e *Entity) Update(dt float32) {

}
