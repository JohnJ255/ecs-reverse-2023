package component

import "github.com/yohamta/donburi"

type PhysicalData struct {
	IsFixed      bool
	Mass         float64
	Friction     float64
	BaseInertion float64
}

var Physical = donburi.NewComponentType[PhysicalData]()
