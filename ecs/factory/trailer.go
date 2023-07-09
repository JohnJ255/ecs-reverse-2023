package factory

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

func NewTrailerForCar(w donburi.World, car *donburi.Entry, look *framework.Sprite) {
	trailer := Trailer.Spawn(w)

	sp := component.Spatial.Get(trailer)
	tr := component.Trailer.Get(trailer)
	sprite := component.Sprite.Get(trailer)
	phys := component.Physical.Get(trailer)
	collider := component.Collider.Get(trailer)

	traktor := component.Traktor.Get(car)
	carSp := component.Spatial.Get(car)

	carEntity := car.Entity()
	trailerEntity := trailer.Entity()
	tr.Traktor = &carEntity
	tr.TowbarUV = framework.VecUV{1, 0.5}
	traktor.Trailer = &trailerEntity

	phys.Mass = 100
	phys.Friction = 1
	phys.IsFixed = false
	phys.BaseInertion = 0.97

	sp.Rotation = carSp.Rotation
	sp.Pivot = framework.VecUV{0.3, 0.5}
	sp.Size = framework.Size{114, 56}

	points := []framework.VecUV{
		{0, 0},
		{0.7, 0},
		{0.7, 0.15},
		{0.95, 0.5},
		{0.7, 0.85},
		{0.7, 1},
		{0, 1},
	}
	cld := framework.NewPolygonColliderUV(points, sp.Size, sp)
	collider.AddFigure(cld)

	look.SetToSize(sp.Size)
	sprite.Angle = look.DrawAngle
	sprite.Image = look.GetSprite()
}
