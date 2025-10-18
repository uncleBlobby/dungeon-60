package tilemap

import (
	"math/rand/v2"

	"github.com/aquilax/go-perlin"
	"github.com/uncleBlobby/dungeon-60/internal/tile"
)

type Tilemap struct {
	Width  int
	Height int
	Tiles  []*tile.Tile
}

func CreateDefault(w, h int) *Tilemap {
	tm := []*tile.Tile{}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			tm = append(tm, tile.Create(x, y, tile.TILE_SIZE, tile.TILETYPE_GRASS))
		}
	}

	r := &Tilemap{
		Width:  w,
		Height: h,
		Tiles:  tm,
	}

	return r
}

func Create(w, h int) *Tilemap {
	tm := []*tile.Tile{}

	alpha := 2.0
	beta := 2.0
	n := int32(3)
	seed := rand.Int64()

	p := perlin.NewPerlin(alpha, beta, n, seed)
	scale := 0.1

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			val := p.Noise2D(float64(x)*scale, float64(y)*scale)

			val = (val + 1) / 2

			if val <= 0.4 {
				tm = append(tm, tile.Create(x, y, tile.TILE_SIZE, tile.TILETYPE_STONE))
			} else if val < 0.8 {
				tm = append(tm, tile.Create(x, y, tile.TILE_SIZE, tile.TILETYPE_GRASS))
			} else {
				tm = append(tm, tile.Create(x, y, tile.TILE_SIZE, tile.TILETYPE_DIRT))
			}
		}
	}

	r := &Tilemap{
		Width:  w,
		Height: h,
		Tiles:  tm,
	}

	return r
}

func (tm *Tilemap) Draw() {
	for _, t := range tm.Tiles {
		t.Draw()
	}
}

func (tm *Tilemap) Update(dt float32) {

}
