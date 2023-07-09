package framework

import "math"

type PolygonLine struct {
	start, end Vec2
}

func (l PolygonLine) IntercectionWithCircle(circle *ColliderShapeCircle) []Vec2 {
	start := l.start.Sub(circle.center)
	end := l.end.Sub(circle.center)
	diff := end.Sub(start)

	a := diff.X*diff.X + diff.Y*diff.Y
	b := 2 * (diff.X*start.X + diff.Y*start.Y)
	c := start.X*start.X + start.Y*start.Y - circle.radius*circle.radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil
	}
	isIntersectInline := func(x float64) bool {
		return x >= 0 && x <= 1
	}

	count := 2
	if discriminant == 0 {
		count = 1
	}
	res := []Vec2{}
	for i := 0; i < count; i++ {
		sign := float64(1 - 2*i)
		x := (-b + sign*math.Sqrt(discriminant)) / (2 * a)
		if isIntersectInline(x) {
			res = append(res, Vec2{l.start.X + x*diff.X, l.start.Y + x*diff.Y})
		}
	}
	return res
}

func (l PolygonLine) IntersectionWithLine(line *PolygonLine) *Vec2 {
	d := (l.end.X-l.start.X)*(line.end.Y-line.start.Y) - (line.end.X-line.start.X)*(l.end.Y-l.start.Y)
	if d == 0 {
		return nil
	}

	lambda := ((l.start.Y-line.start.Y)*(line.end.X-line.start.X) - (l.start.X-line.start.X)*(line.end.Y-line.start.Y) + 1) / d
	gamma := ((l.start.Y-line.start.Y)*(l.end.X-l.start.X) - (l.start.X-line.start.X)*(l.end.Y-l.start.Y) + 1) / d

	if (0 < lambda && lambda < 1) && (0 < gamma && gamma < 1) {
		dx := l.end.X - l.start.X
		dy := l.end.Y - l.start.Y
		// todo: test it
		return &Vec2{dx*lambda + l.start.X, dy*lambda + l.start.Y}
	}

	return nil
}

func (l PolygonLine) Normal() Vec2 {
	v := l.ToVec2()
	return Vec2{-v.Y, v.X}.Normalize()
}

func (l PolygonLine) ToVec2() Vec2 {
	return l.end.Sub(l.start)
}
