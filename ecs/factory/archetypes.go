package factory

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"github.com/yohamta/donburi"
)

var (
	Wall = newArchetype(
		component.Physical,
		component.Spatial,
		component.Collider,
		tags.GameElement,
	)
	Trigger = newArchetype(
		component.Spatial,
		component.Trigger,
		component.Collider,
		tags.GameElement,
	)
	Car = newArchetype(
		component.Car,
		component.Traktor,
		component.Spatial,
		component.Sprite,
		component.Collider,
		component.Physical,
		tags.GameElement,
	)
	Trailer = newArchetype(
		component.Trailer,
		component.Spatial,
		component.Collider,
		component.Sprite,
		component.Physical,
		tags.GameElement,
	)
)

type archetype struct {
	components []donburi.IComponentType
}

func newArchetype(cs ...donburi.IComponentType) *archetype {
	return &archetype{
		components: cs,
	}
}

func (a archetype) Spawn(w donburi.World) *donburi.Entry {
	return w.Entry(w.Create(a.components...))
}
