package factory

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"ecs_test_cars/helper"
	"github.com/yohamta/donburi"
)

func NewPlayerCar(w donburi.World, x, y float64, angle framework.Radian, sprite *framework.Sprite) *donburi.Entry {
	entry, _, _ := NewCar(w, x, y, angle, sprite)
	entry.AddComponent(tags.Player)
	entry.AddComponent(tags.InScreen)

	return entry
}

func NewCar(w donburi.World, x, y float64, r framework.Radian, look *framework.Sprite) (*donburi.Entry, *component.CarData, *component.SpatialData) {
	entry := Car.Spawn(w)

	car := component.Car.Get(entry)
	phys := component.Physical.Get(entry)
	sp := component.Spatial.Get(entry)
	sprite := component.Sprite.Get(entry)
	collider := component.Collider.Get(entry)
	traktor := component.Traktor.Get(entry)

	car.Powerful = 160
	car.BackMaxSpeed = -helper.KmphToPixelsPerTick(54)
	car.MaxSpeed = helper.KmphToPixelsPerTick(180)
	car.Handling = 0.02
	car.SpeedHandling = 0.7
	car.MaxWheelAngle = framework.Degrees(45).ToRadians()
	car.WheelBase = 80

	phys.Mass = 800
	phys.Friction = 1
	phys.IsFixed = false
	phys.BaseInertion = 0.97

	traktor.TowbarUV = framework.VecUV{0, 0.5}

	sp.Position.X = x
	sp.Position.Y = y
	sp.Rotation = r
	sp.Size = framework.Size{114, 56}
	sp.Pivot = framework.VecUV{0.2, 0.5}

	collider.AddFigure(framework.NewBoxCollider(sp.Size, sp))

	look.SetToSize(sp.Size)
	sprite.Angle = look.DrawAngle
	sprite.Image = look.GetSprite()

	return entry, car, sp
}
