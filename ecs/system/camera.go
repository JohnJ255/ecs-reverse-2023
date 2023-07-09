package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type Camera struct {
	f *framework.Framework
}

func NewFollowCamera(f *framework.Framework) *Camera {
	return &Camera{
		f: f,
	}
}

func (c *Camera) Update(ecs *ecs.ECS) {
	q := donburi.NewQuery(filter.Contains(tags.Player, component.Spatial))
	player, ok := q.First(ecs.World)
	if !ok {
		return
	}
	level, ok := component.CurrentLevel.First(ecs.World)
	if !ok {
		return
	}
	levelData := component.CurrentLevel.Get(level)
	f := levelData.LevelFiller
	if f == nil {
		return
	}
	sp := component.Spatial.Get(player)
	pos := sp.Position.Sub(framework.Vec2{X: 300, Y: 300})
	w := float64(c.f.WindowWidth)
	h := float64(c.f.WindowHeight)
	size := f.GetSize().Sub(framework.Vec2{X: w, Y: h})
	pos.X = framework.Limited(pos.X, 0, size.Length)
	pos.Y = framework.Limited(pos.Y, 0, size.Height)
	levelData.Offset = pos.Mul(-1)
}
