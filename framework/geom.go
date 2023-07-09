package framework

import "math"

type Size struct {
	Length, Height float64
}

func (s Size) AsVec2() Vec2 {
	return Vec2{s.Height, s.Length}
}

func (s Size) Sub(size Vec2) Size {
	s.Height -= size.Y
	s.Length -= size.X
	return s
}

type IntSize struct {
	Width, Height int
}

func (s IntSize) AsVec2() Vec2 {
	return Vec2{float64(s.Width), float64(s.Height)}
}

type DirectionPosition struct {
	Vec2
	Angle Radian
}

type Vec2 struct {
	X, Y float64
}

type Bounds struct {
	Min, Max Vec2
}

type VecUV struct {
	U, V float64
}

func (u VecUV) ToVec2(s Size) Vec2 {
	return Vec2{u.U * s.Length, u.V * s.Height}
}

type Projection struct {
	Min, Max float64
}

type Degrees float64
type Radian float64

var AngleTop = Radian(-math.Pi / 2)
var AngleBottom = Radian(math.Pi / 2)
var AngleLeft = Radian(math.Pi)
var AngleRight = Radian(0)

func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func NewDPos(x, y float64, a Radian) DirectionPosition {
	return DirectionPosition{Vec2{X: x, Y: y}, a}
}

func (d Degrees) ToRadians() Radian {
	return Radian(float64(d) * math.Pi / 180)
}

func (r Radian) ToVec() Vec2 {
	return Vec2{math.Cos(float64(r)), math.Sin(float64(r))}
}

func (r Radian) ToDegrees() Degrees {
	return Degrees(float64(r) * 180 / math.Pi)
}

func (r Radian) LefterThan(r2 Radian) bool {
	_ = r.ToDegrees()
	_ = r2.ToDegrees()
	if r == r2 {
		return false
	}
	r = r.Normalize()
	r2 = r2.Normalize()
	if math.Abs(float64(r-r2)) < 0.00001 {
		return false
	}
	a := float64((r2 - r).Normalize())
	b := float64(((r + 2*math.Pi) - r2).Normalize())
	return math.Abs(a) > math.Abs(b)
}

func (r Radian) RighterThan(r2 Radian) bool {
	if r == r2 {
		return false
	}
	//fc := Radian(2 * math.Pi)
	//for r < 0 {
	//	r += fc
	//}
	//for r2 < 0 {
	//	r2 += fc
	//}
	//for r >= fc {
	//	r -= fc
	//}
	//for r2 >= fc {
	//	r2 -= fc
	//}
	return r2-r > math.Pi || r2-r < 0
}

func (r Radian) Normalize() Radian {
	for r < 0 {
		r += 2 * math.Pi
	}
	for r >= 2*math.Pi {
		r -= 2 * math.Pi
	}
	return r
}

func (r Radian) NormalizePi2() Radian {
	for r < -math.Pi/2 {
		r += 2 * math.Pi
	}
	for r >= 3*math.Pi/2 {
		r -= 2 * math.Pi
	}
	return r
}

func (r Radian) F64() float64 {
	return float64(r)
}

func (r Radian) Cos() float64 {
	return math.Cos(float64(r))
}

func (r Radian) Sin() float64 {
	return math.Sin(float64(r))
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Mul(n float64) Vec2 {
	return Vec2{v.X * n, v.Y * n}
}

func (v Vec2) ScalarMul(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) MulXY(x float64, y float64) Vec2 {
	return Vec2{v.X * x, v.Y * y}
}

func (v Vec2) Div(f float64) Vec2 {
	return Vec2{v.X / f, v.Y / f}
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	if v.X == 0 && v.Y == 0 {
		return v
	}
	return v.Mul(1 / v.Length())
}

func (v Vec2) ToRadian() Radian {
	return Radian(math.Atan2(v.Y, v.X))
}

func (v Vec2) Rotate(angle Radian) Vec2 {
	l := v.Length()
	a := v.ToRadian() + angle
	return a.ToVec().Mul(l)
}

func (v Vec2) RotateAround(angle Radian, b Vec2) Vec2 {
	return v.Sub(b).Rotate(angle).Add(b)
}

func (v Vec2) InRectWithoutRotation(pos Vec2, size Size) bool {
	v = v.Sub(pos)
	return v.X >= 0 && v.X <= size.Length && v.Y >= 0 && v.Y <= size.Height
}

func CalcCenter(points []Vec2) *Vec2 {
	if len(points) == 0 {
		return nil
	}
	res := Vec2{0, 0}
	for _, v := range points {
		res = res.Add(v)
	}
	res = res.Div(float64(len(points)))
	return &res
}

func (p *Projection) Cross(p2 Projection) float64 {
	return math.Min(p.Max, p2.Max) - math.Max(p.Min, p2.Min)
}

func (p *Projection) Into(p2 Projection) bool {
	return p.Min >= p2.Min && p.Max <= p2.Max
}
