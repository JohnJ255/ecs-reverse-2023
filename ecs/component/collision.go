package component

import (
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

type CollisionData struct {
	Contacts   []framework.ContactSet
	WithEntity donburi.Entity
}

var Collision = donburi.NewComponentType[CollisionData]()
