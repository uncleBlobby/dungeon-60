package systems

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/uncleBlobby/dungeon-60/internal/entity"
	"github.com/uncleBlobby/dungeon-60/internal/world"
)

type SpriteSystem struct {
	World    world.World
	Textures map[string]rl.Texture2D
}

func InitTextures(w world.World) *SpriteSystem {

	ss := &SpriteSystem{
		Textures: map[string]rl.Texture2D{},
		World:    w,
	}
	ss.Textures[string(entity.ENTITY_BUSH)] = rl.LoadTexture("assets/sprites/bush1.png")
	ss.Textures[string(entity.ENTITY_TREE)] = rl.LoadTexture("assets/sprites/tree1.png")

	return ss
}

func (ss *SpriteSystem) InitSprites() {
	ents := ss.World.GetEntities()
	for i := 0; i < len(ents); i++ {
		// load sprites based on entity type?
		fmt.Println(ents[i])

		switch ents[i].Type {
		case entity.ENTITY_BUSH:
			ents[i].Sprite = ss.Textures[string(entity.ENTITY_BUSH)]

		case entity.ENTITY_TREE:
			ents[i].Sprite = ss.Textures[string(entity.ENTITY_TREE)]
		}
	}
}
