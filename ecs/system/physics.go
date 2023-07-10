package system

import (
	"ecs_test_cars/ecs/common"
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type Physics struct {
	physicsQuery *donburi.Query
}

func NewPhysics() *Physics {
	return &Physics{
		physicsQuery: donburi.NewQuery(filter.Contains(component.Physical, component.Spatial, component.Collision)),
	}
}

func (p *Physics) Update(ecs *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(ecs.World); ok {
		return
	}

	p.physicsQuery.Each(ecs.World, func(entry *donburi.Entry) {
		collide := *component.Collision.Get(entry)
		other := ecs.World.Entry(collide.WithEntity)
		entry.RemoveComponent(component.Collision)
		if !other.HasComponent(component.Physical) {
			return
		}
		otherPhys := component.Physical.Get(other)
		if otherPhys == nil {
			return
		}
		phys := component.Physical.Get(entry)
		if phys.IsFixed && otherPhys.IsFixed {
			return
		}
		sp := component.Spatial.Get(entry)

		massesK := otherPhys.Mass / phys.Mass

		isContactWithOwnTraktor := false
		isContactWithSelfTrailer := false
		isTrailer := entry.HasComponent(component.Trailer)
		if isTrailer {
			withTraktor := other.HasComponent(component.Traktor)
			if withTraktor {
				traktor := component.Traktor.Get(other)
				isContactWithOwnTraktor = traktor.Trailer != nil && *traktor.Trailer == entry.Entity()
			}
		} else {
			isTraktor := entry.HasComponent(component.Traktor)
			if isTraktor {
				withTrailer := other.HasComponent(component.Trailer)
				if withTrailer {
					trailer := component.Trailer.Get(other)
					isContactWithSelfTrailer = trailer.Traktor != nil && *trailer.Traktor == entry.Entity()
				}
			}
		}

		for _, cs := range collide.Contacts {
			if isContactWithOwnTraktor {
				p.traktorPhysics(entry, other, cs)
			} else if isContactWithSelfTrailer {
			} else {
				p.basePhysics(sp, cs, massesK)
				if isTrailer {
					p.traileredTraktorPhysics(ecs.World, entry, cs)
				}
			}
		}
	})
}

func (p *Physics) calcAngleStep(a float64) framework.Radian {
	return framework.Radian(0.01 * a)
}

func (p *Physics) basePhysics(sp *component.SpatialData, cs framework.ContactSet, massesK float64) {
	sp.Position.X += cs.MoveOut.X
	sp.Position.Y += cs.MoveOut.Y

	angle := sp.Position.Sub(*cs.Center).ToRadian()

	sign := -p.calcAngleStep(massesK)
	if angle.LefterThan((*cs.MoveOut).ToRadian()) {
		sign *= -1
	}
	k := 0.01
	if massesK < 1 {
		k = 3
	}
	sp.Rotation += sign * framework.Radian(k)
}

func (p *Physics) traktorPhysics(trailer *donburi.Entry, traktor *donburi.Entry, cs framework.ContactSet) {
	trailerSp := component.Spatial.Get(trailer)
	traktorSp := component.Spatial.Get(traktor)
	trailerTowbar := component.Trailer.Get(trailer)
	traktorTowbar := component.Traktor.Get(traktor)

	sign := framework.Radian(0.22)
	if cs.MoveOut.ToRadian().LefterThan(trailerSp.Rotation) {
		sign = -sign
	}

	trailerSp.Rotation += sign

	tlp := common.GetTowbarLocalPosition(trailerSp, trailerTowbar.TowbarUV)
	traktorTowbarPos := common.GetTowbarPosition(traktorSp, traktorTowbar.TowbarUV)
	trailerSp.Position = traktorTowbarPos.Sub(tlp)
}

func (p *Physics) traileredTraktorPhysics(w donburi.World, trailerEntry *donburi.Entry, cs framework.ContactSet) {
	trailer := component.Trailer.Get(trailerEntry)
	if trailer.Traktor != nil {
		k := cs.MoveOut.Length()
		sp := component.Spatial.Get(w.Entry(*trailer.Traktor))
		car := component.Car.Get(w.Entry(*trailer.Traktor))
		sp.Position.X += (k - car.Speed) * sp.Rotation.Cos()
		sp.Position.Y += (k - car.Speed) * sp.Rotation.Sin()
	}
}
