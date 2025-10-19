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
	fmt.Println("LOADING BUSH TEXTURE")
	ss.Textures[string(entity.ENTITY_BUSH)] = rl.LoadTexture("assets/sprites/bush1.png")
	fmt.Println("LOADING TREE TEXTURE")
	ss.Textures[string(entity.ENTITY_TREE)] = rl.LoadTexture("assets/sprites/tree1.png")

	return ss
}

func (ss *SpriteSystem) InitSprites() {
	fmt.Println("INIT SPRITES")
	ents := ss.World.GetEntities()
	fmt.Println("GOT ENTITIES")
	fmt.Println(ents)
	for i := 0; i < len(ents); i++ {
		// load sprites based on entity type?
		fmt.Println(ents[i])

		switch ents[i].Type {
		case entity.ENTITY_BUSH:
			// load bush sprite
			ents[i].Sprite = ss.Textures[string(entity.ENTITY_BUSH)]

		case entity.ENTITY_TREE:
			ents[i].Sprite = ss.Textures[string(entity.ENTITY_TREE)]
		}
	}
}
