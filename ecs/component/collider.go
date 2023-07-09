package component

import (
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

var Collider = donburi.NewComponentType[framework.Collider]()
