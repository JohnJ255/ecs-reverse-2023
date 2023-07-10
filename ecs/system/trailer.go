package system

import (
	"ecs_test_cars/ecs/common"
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type TrailerMoving struct {
}

func (tm *TrailerMoving) Update(e *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(e.World); ok {
		return
	}

	donburi.NewQuery(filter.Contains(component.Trailer, component.Spatial, component.Physical)).Each(e.World, func(entry *donburi.Entry) {
		t := component.Trailer.Get(entry)

		if t.Traktor == nil {
			return
		}
		traktor := e.World.Entry(*t.Traktor)

		tm.followTraktor(traktor, entry)
	})
}

func (tm *TrailerMoving) followTraktor(traktor *donburi.Entry, trailer *donburi.Entry) {
	trailerSp := component.Spatial.Get(trailer)
	trailerData := component.Trailer.Get(trailer)
	traktorSp := component.Spatial.Get(traktor)
	traktorTowbar := component.Traktor.Get(traktor)

	traktorTowbarPos := common.GetTowbarPosition(traktorSp, traktorTowbar.TowbarUV)
	if traktorTowbar.TraktorJointPosition == nil {
		traktorTowbar.TraktorJointPosition = &traktorTowbarPos
	}
	velocity := traktorTowbarPos.Sub(*traktorTowbar.TraktorJointPosition)

	tlp := common.GetTowbarLocalPosition(trailerSp, trailerData.TowbarUV)

	trailerSp.Position = traktorTowbarPos.Sub(tlp)
	trailerSp.Rotation = tlp.Add(velocity).ToRadian()
	traktorTowbar.TraktorJointPosition = &traktorTowbarPos
}
