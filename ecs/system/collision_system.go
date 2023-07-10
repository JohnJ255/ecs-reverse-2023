package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type Collision struct {
	f *framework.Framework
}

func NewCollisionSystem() *Collision {
	return &Collision{}
}

func (c *Collision) Update(ecs *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(ecs.World); ok {
		return
	}

	component.Collider.Each(ecs.World, func(entry *donburi.Entry) {
		collision := component.Collider.Get(entry)

		for _, otherCollisionEntry := range c.getClosestCollisionsFor(entry.Entity(), ecs.World) {
			collides := make([]*framework.Collide, 0)
			cs := collision.Intersect(component.Collider.Get(otherCollisionEntry))
			if len(cs) > 0 && cs[0].MoveOut != nil {
				collides = append(collides, framework.NewCollide(collision, cs))
				if !entry.HasComponent(component.Collision) {
					entry.AddComponent(component.Collision)
					contacts := component.Collision.Get(entry)
					contacts.WithEntity = otherCollisionEntry.Entity()
					contacts.Contacts = make([]framework.ContactSet, 0)
					for _, collide := range collides {
						contacts.Contacts = append(contacts.Contacts, collide.Contacts...)
					}
				}
			}
		}
	})
}

func (c *Collision) getClosestCollisionsFor(entity donburi.Entity, w donburi.World) []*donburi.Entry {
	res := make([]*donburi.Entry, 0)
	component.Collider.Each(w, func(entry *donburi.Entry) {
		if entry.Entity() == entity {
			return
		}
		res = append(res, entry)
	})

	return res
}
