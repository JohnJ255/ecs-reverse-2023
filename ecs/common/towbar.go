package common

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/framework"
	"math"
)

func GetTowbarLocalPosition(sp *component.SpatialData, towbarUV framework.VecUV) framework.Vec2 {
	dx := sp.Size.Length * (towbarUV.U - sp.Pivot.U)
	dy := sp.Size.Height * (towbarUV.V - sp.Pivot.V)
	length := math.Sqrt(dx*dx + dy*dy)
	angle := sp.Rotation + framework.Radian(math.Atan2(dy, dx))
	x := length * angle.Cos()
	y := length * angle.Sin()

	return framework.Vec2{x, y}
}

func GetTowbarPosition(sp *component.SpatialData, towbarUV framework.VecUV) framework.Vec2 {
	return sp.Position.Add(GetTowbarLocalPosition(sp, towbarUV))
}
