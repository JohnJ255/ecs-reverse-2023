package component

import (
	"ecs_test_cars/framework"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type ILevel interface {
	Fill(ecs donburi.World)
	GetSize() framework.Size
	GetSprite() *ebiten.Image
}

type CurrentLevelData struct {
	Index       int
	IsFilled    bool
	LevelFiller ILevel
	Name        string
	Offset      framework.Vec2
}

var CurrentLevel = donburi.NewComponentType[CurrentLevelData]()
