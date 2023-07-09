package component

import (
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

type TraktorData struct {
	Trailer              *donburi.Entity
	TowbarUV             framework.VecUV
	TraktorJointPosition *framework.Vec2
}

var Traktor = donburi.NewComponentType[TraktorData]()
