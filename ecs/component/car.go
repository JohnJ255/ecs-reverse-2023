package component

import (
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
)

type CarData struct {
	WheelAngle    framework.Radian
	MaxWheelAngle framework.Radian
	WheelBase     float64
	Speed         float64
	Powerful      float64
	BackMaxSpeed  float64
	MaxSpeed      float64
	Handling      float64
	SpeedHandling float64

	Accelerate  float64
	WheelRotate float64
}

var Car = donburi.NewComponentType[CarData]()
