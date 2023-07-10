package component

import (
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

type MenuData struct {
	Elements []framework.IGUIElement
	Position framework.Vec2
	Size     framework.Vec2
	IsOpened bool
	Caption  string
}

var Menu = donburi.NewComponentType[MenuData]()
