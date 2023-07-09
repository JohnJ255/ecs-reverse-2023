package component

import (
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

type TrailerType int

const (
	TrailerTypeNone TrailerType = iota
	TrailerTypeCart
	TrailerTypeTrailer
)

type TrailerData struct {
	Traktor  *donburi.Entity
	TowbarUV framework.VecUV
}

var Trailer = donburi.NewComponentType[TrailerData]()
