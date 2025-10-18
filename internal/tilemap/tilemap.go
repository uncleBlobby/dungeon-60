package tilemap

import "github.com/uncleBlobby/dungeon-60/internal/tile"

type Tilemap struct {
	Width  int
	Height int
	Tiles  []*tile.Tile
}

func Create(w, h int) *Tilemap {
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

func (tm *Tilemap) Draw() {
	for _, t := range tm.Tiles {
		t.Draw()
	}
}

func (tm *Tilemap) Update(dt float32) {

}
