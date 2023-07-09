package framework

import (
	"math"
	"sort"
)

type ColliderShapePolygon struct {
	points []Vec2
	owner  ICollideOwner
}

func (p *ColliderShapePolygon) Intersect(other IColliderFigure) ContactSet {
	contactSet := ContactSet{}
	switch other.(type) {
	case *ColliderShapeCircle:
		circle := other.(*ColliderShapeCircle)
		for _, line := range p.GetRealLines() {
			contactSet.Points = append(contactSet.Points, line.IntercectionWithCircle(circle)...)
		}
	case *ColliderShapePolygon:
		polygon := other.(*ColliderShapePolygon)
		for _, line := range p.GetRealLines() {
			for _, otherLine := range polygon.GetRealLines() {
				if point := line.IntersectionWithLine(otherLine); point != nil {
					contactSet.Points = append(contactSet.Points, *point)
				}
			}
		}
	}

	if len(contactSet.Points) == 0 {
		return contactSet
	}

	contactSet.Center = CalcCenter(contactSet.Points)
	contactSet.MoveOut = p.CalcMoveOut(contactSet, other)

	return contactSet
}

func (c *ColliderShapePolygon) SetOffset(offset Vec2) {
	for i, p := range c.points {
		c.points[i].X = p.X + offset.X
		c.points[i].Y = p.Y + offset.Y
	}
}

func (c *ColliderShapePolygon) SetScale(scale Vec2) {
	// todo: test it
	center := CalcCenter(c.points)
	if center == nil {
		return
	}
	for i, p := range c.points {
		np := center.Sub(p)
		length := np.Length()
		p = np.Mul(scale.X * length)
		c.points[i].X = p.X
		c.points[i].Y = p.Y
	}
}

func (c *ColliderShapePolygon) SetRotation(rot Radian) {
	// todo: test it
	center := CalcCenter(c.points)
	if center == nil {
		return
	}
	for i, p := range c.points {
		p = p.RotateAround(rot, *center)
		c.points[i].X = p.X
		c.points[i].Y = p.Y
	}
}

func (c *ColliderShapePolygon) Bounds() Bounds {
	if len(c.points) < 2 {
		return Bounds{}
	}
	min := c.points[0]
	max := c.points[0]
	for _, p := range c.points {
		if p.X < min.X {
			min.X = p.X
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}

	return Bounds{min, max}
}

func (p *ColliderShapePolygon) crossPolySAT(other *ColliderShapePolygon, axes []Vec2) *Vec2 {
	smallest := Vec2{math.MaxFloat64, 0}

	for _, axis := range axes {
		pr1 := p.Projection(axis)
		pr2 := other.Projection(axis)

		cross := pr1.Cross(pr2)

		if cross <= 0 {
			return nil
		}
		if smallest.Length() > cross {
			if pr1.Min < pr2.Min {
				smallest = axis.Mul(-cross)
			} else {
				smallest = axis.Mul(cross)
			}

		}
	}
	return &smallest
}

func (p *ColliderShapePolygon) CalcMoveOut(set ContactSet, other IColliderFigure) *Vec2 {
	res := Vec2{}
	switch otherP := other.(type) {
	case *ColliderShapeCircle:
		if set.Center == nil {
			return nil
		}
		verts := p.getRealPoints()
		verts = append(verts, *set.Center)
		sort.Slice(verts, func(i, j int) bool {
			return verts[i].Sub(otherP.center).Length() < verts[j].Sub(otherP.center).Length()
		})
		res = Vec2{otherP.center.X - verts[0].X, otherP.center.Y - verts[0].Y}
		res = res.Normalize().Mul(res.Length() - otherP.radius)

		return &res

	case *ColliderShapePolygon:
		v1 := p.crossPolySAT(otherP, p.axesSAT())
		v2 := p.crossPolySAT(otherP, otherP.axesSAT())
		if v1 == nil || v2 == nil {
			return nil
		}
		if v1.Length() < v2.Length() {
			return v1
		}
		return v2
	}

	return nil
}

func (p *ColliderShapePolygon) axesSAT() []Vec2 {
	lines := p.GetRealLines()
	axes := make([]Vec2, 0, len(lines))
	for _, line := range lines {
		axes = append(axes, line.Normal())
	}
	return axes
}

func (p *ColliderShapePolygon) Projection(axis Vec2) Projection {
	axis = axis.Normalize()
	points := p.getRealPoints()
	min := axis.ScalarMul(points[0])
	max := min
	for i := 1; i < len(points); i++ {
		sm := axis.ScalarMul(points[i])
		if sm < min {
			min = sm
		} else if sm > max {
			max = sm
		}
	}
	return Projection{Min: min, Max: max}
}

func (p *ColliderShapePolygon) getRealPoints() []Vec2 {
	points := make([]Vec2, len(p.points))
	for i, point := range p.points {
		points[i] = point.Add(p.owner.GetPosition())
	}
	center := p.owner.GetPosition()
	scale := p.owner.GetScale()
	for i, point := range points {
		points[i] = point.Sub(center).MulXY(scale.X, scale.Y).Add(center)
	}
	rot := p.owner.GetRotation()
	for i, point := range points {
		points[i] = point.RotateAround(rot, center)
	}
	return points
}

func (p *ColliderShapePolygon) GetOwner() ICollideOwner {
	return p.owner
}

func (p *ColliderShapePolygon) getLines(points []Vec2) []*PolygonLine {
	lines := make([]*PolygonLine, 0)
	for i := 0; i < len(points); i++ {
		next := i + 1
		if i == len(points)-1 {
			next = 0
		}
		lines = append(lines, &PolygonLine{points[i], points[next]})
	}

	return lines
}

func (p *ColliderShapePolygon) GetRealLines() []*PolygonLine {
	return p.getLines(p.getRealPoints())
}
