package systems

import "github.com/uncleBlobby/dungeon-60/internal/world"

type EntitySystem struct {
	World world.World
}

func (es *EntitySystem) Update(dt float32) {
	for _, e := range es.World.GetEntities() {
		e.Update(dt)
	}
}

func (es *EntitySystem) Draw() {
	for _, e := range es.World.GetEntities() {
		e.Draw()
	}
}
