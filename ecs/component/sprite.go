package component

import (
	"ecs_test_cars/framework"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image *ebiten.Image
	Angle framework.Radian
}

var Sprite = donburi.NewComponentType[SpriteData]()
