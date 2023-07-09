package levels

import (
	"ecs_test_cars/framework"
	"ecs_test_cars/loader"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type LevelAbout struct {
}

func (l *LevelAbout) GetSize() framework.Size {
	return framework.Size{800, 600}
}

func (l *LevelAbout) GetSprite() *ebiten.Image {
	sprite := framework.InitSprites(l.GetSize())
	sprite.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[7])
	return sprite.GetSprite()
}

func (l *LevelAbout) Fill(w donburi.World) {
}
