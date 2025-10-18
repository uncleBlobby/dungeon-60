package world

import (
	"github.com/uncleBlobby/dungeon-60/internal/entity"
	"github.com/uncleBlobby/dungeon-60/internal/level"
	"github.com/uncleBlobby/dungeon-60/internal/player"
)

type World interface {
	GetLevel() *level.Level
	GetPlayer() *player.Player
	GetEntities() []*entity.Entity
}
