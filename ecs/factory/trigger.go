package factory

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

func NewTrigger(w donburi.World, t component.TriggerType, pos framework.Vec2, size framework.Size) *donburi.Entry {
	entry := Trigger.Spawn(w)
	component.Trigger.SetValue(entry, component.TriggerData{
		TriggerType: t,
	})
	sp := component.SpatialData{
		Position: pos,
		Size:     size,
	}
	component.Spatial.SetValue(entry, sp)
	component.Collider.SetValue(entry, *framework.InitCollider(framework.NewBoxCollider(size, &sp)))

	return entry
}
