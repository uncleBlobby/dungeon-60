package systems

import "github.com/uncleBlobby/dungeon-60/internal/world"

type PlayerSystem struct {
	World world.World
}

func (ps *PlayerSystem) Draw() {
	p := ps.World.GetPlayer()
	p.Draw()
}

func (ps *PlayerSystem) Update(dt float32) {
	p := ps.World.GetPlayer()
	p.Update(dt)
}
