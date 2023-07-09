package component

import (
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

type SpatialData struct {
	Position framework.Vec2
	Size     framework.Size
	Rotation framework.Radian
	Pivot    framework.VecUV
}

func (s *SpatialData) GetPosition() framework.Vec2 {
	return s.Position
}

func (s *SpatialData) GetRotation() framework.Radian {
	return s.Rotation
}

func (s *SpatialData) SetPosition(pos framework.Vec2) {
	s.Position = pos
}

func (s SpatialData) SetRotation(rot framework.Radian) {
	s.Rotation = rot
}

func (s SpatialData) GetScale() framework.Vec2 {
	return framework.Vec2{1, 1}
}

func (s SpatialData) GetPivot() framework.VecUV {
	return s.Pivot
}

var Spatial = donburi.NewComponentType[SpatialData]()
