package systems

import "github.com/uncleBlobby/dungeon-60/internal/world"

type LevelSystem struct {
	World world.World
}

func (ls *LevelSystem) Draw() {
	level := ls.World.GetLevel()
	level.Draw()
}

func (ls *LevelSystem) Update(dt float32) {
	level := ls.World.GetLevel()
	level.Update(dt)
}
