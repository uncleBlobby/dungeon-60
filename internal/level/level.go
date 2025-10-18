package level

import "github.com/uncleBlobby/dungeon-60/internal/tilemap"

type Level struct {
	Width   int
	Height  int
	Tilemap *tilemap.Tilemap
}

func (l *Level) Draw() {
	l.Tilemap.Draw()
}

func Create(w, h int) *Level {

	l := &Level{
		Width:   w,
		Height:  h,
		Tilemap: tilemap.Create(w, h),
	}

	return l
}

func (l *Level) Update(dt float32) {

}
