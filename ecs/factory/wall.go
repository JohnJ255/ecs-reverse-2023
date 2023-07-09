package factory

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

func NewWall(w donburi.World, pos framework.Vec2, size framework.Size) *donburi.Entry {
	wall := Wall.Spawn(w)
	sp := component.SpatialData{
		Position: pos,
		Size:     size,
	}
	component.Spatial.SetValue(wall, sp)
	component.Collider.SetValue(wall, *framework.InitCollider(framework.NewBoxCollider(size, &sp)))
	component.Physical.SetValue(wall, component.PhysicalData{
		IsFixed:      true,
		Mass:         1000,
		Friction:     1,
		BaseInertion: 0,
	})
	return wall
}
