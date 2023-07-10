package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type InScreenLimit struct {
	f     *framework.Framework
	limit framework.Size
}

func NewScreenLimiter(f *framework.Framework) *InScreenLimit {
	return &InScreenLimit{
		f: f,
	}
}

func (isl *InScreenLimit) Update(e *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(e.World); ok {
		return
	}

	donburi.NewQuery(filter.Contains(component.Spatial, tags.InScreen)).Each(e.World, func(entry *donburi.Entry) {
		levelEntry, ok := component.CurrentLevel.First(e.World)
		if !ok {
			return
		}
		level := component.CurrentLevel.Get(levelEntry)
		if level.LevelFiller == nil {
			return
		}
		size := level.LevelFiller.GetSize()
		sp := component.Spatial.Get(entry)
		w := float64(isl.f.WindowWidth)
		h := float64(isl.f.WindowHeight)
		isl.limit = size.Sub(framework.Vec2{X: w, Y: h})
		sp.Position.X = framework.Limited(sp.Position.X, 0, w+isl.limit.Length)
		sp.Position.Y = framework.Limited(sp.Position.Y, 0, h+isl.limit.Height)
	})
}
