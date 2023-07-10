package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"math"
)

const massEtalon = 1000

type CarMoving struct {
}

func (cm *CarMoving) Update(e *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(e.World); ok {
		return
	}

	donburi.NewQuery(filter.Contains(component.Car, component.Spatial, component.Physical)).Each(e.World, func(entry *donburi.Entry) {
		c := component.Car.Get(entry)
		s := component.Spatial.Get(entry)
		p := component.Physical.Get(entry)

		cm.control(c, s, p)
	})
}

func (cm *CarMoving) control(c *component.CarData, s *component.SpatialData, p *component.PhysicalData) {
	powerful := c.Powerful / p.Mass
	k := 1 + (massEtalon-p.Mass)/massEtalon
	minSpeed := c.BackMaxSpeed * k
	maxSpeed := c.MaxSpeed * k
	inertion := cm.calcInertionDependsMass(p)
	if c.Accelerate == 0 && c.Speed != 0 {
		c.Speed *= inertion
		if math.Abs(c.Speed) < powerful {
			c.Speed = 0
		}
	} else {
		c.Speed = framework.Limited(c.Speed+c.Accelerate*powerful, minSpeed, maxSpeed)
	}
	maxWheelAngle := framework.Radian(float64(c.MaxWheelAngle) * (maxSpeed - math.Abs(c.Speed)*c.SpeedHandling) / maxSpeed)
	newWheelAngle := framework.Radian(framework.Stepped(float64(c.WheelAngle), c.WheelRotate*float64(maxWheelAngle), c.Handling))
	c.WheelAngle = framework.Limited(newWheelAngle, -maxWheelAngle, maxWheelAngle)

	s.Rotation += framework.Radian(math.Atan2(c.WheelBase*math.Tan(float64(c.WheelAngle)), s.Size.Length+c.WheelBase) * c.Speed * 0.03)
	dx := c.Speed * math.Cos(s.Rotation.F64())
	dy := c.Speed * math.Sin(s.Rotation.F64())
	s.Position.X += dx
	s.Position.Y += dy
}

func (cm *CarMoving) calcInertionDependsMass(p *component.PhysicalData) float64 {
	mass := p.Mass
	//if c.Trailer != nil {
	//	mass += c.Trailer.getSelfMass()
	//}
	k := 1 + (massEtalon-mass)/massEtalon
	return framework.Limited(p.BaseInertion-k/10, 0.9, 0.999)
}
