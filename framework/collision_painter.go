package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

type DefaultColliderPainter struct {
	color color.Color
}

func (painter *DefaultColliderPainter) Draw(dst *ebiten.Image, figure IColliderFigure) {
	switch figure.(type) {
	case *ColliderShapeCircle:
		circle := figure.(*ColliderShapeCircle)
		center := circle.GetCenter()
		vector.StrokeCircle(dst, float32(center.X), float32(center.Y), float32(circle.GetRadius()), 1, painter.color, false)
	case *ColliderShapePolygon:
		polygon := figure.(*ColliderShapePolygon)
		for _, line := range polygon.GetRealLines() {
			vector.StrokeLine(dst, float32(line.start.X), float32(line.start.Y), float32(line.end.X), float32(line.end.Y), 1, painter.color, false)
		}
	}
}

type DefaultIntersectionPainter struct {
	color       color.Color
	arrowColor  color.Color
	arrowLength float64
	arrowAngle  Radian
}

func (painter *DefaultIntersectionPainter) Draw(dst *ebiten.Image, cs ContactSet) {
	for i := 0; i < len(cs.Points)-1; i++ {
		p1 := cs.Points[i]
		p2 := cs.Points[i+1]
		vector.StrokeLine(dst, float32(p1.X), float32(p1.Y), float32(p2.X), float32(p2.Y), 2, painter.color, false)
	}
	if cs.MoveOut != nil && cs.Center != nil {
		painter.DrawArrow(dst, *cs.Center, (*cs.Center).Add(*cs.MoveOut))
	}
}

func (painter *DefaultIntersectionPainter) DrawArrow(dst *ebiten.Image, from Vec2, to Vec2) {
	vector.StrokeLine(dst, float32(from.X), float32(from.Y), float32(to.X), float32(to.Y), 1, painter.arrowColor, false)

	aran := 3*math.Pi/4 + painter.arrowAngle
	left := (to.Sub(from).ToRadian() - aran).ToVec().Mul(painter.arrowLength).Add(to)
	right := (to.Sub(from).ToRadian() + aran).ToVec().Mul(painter.arrowLength).Add(to)

	vector.StrokeLine(dst, float32(to.X), float32(to.Y), float32(left.X), float32(left.Y), 1, painter.arrowColor, false)
	vector.StrokeLine(dst, float32(to.X), float32(to.Y), float32(right.X), float32(right.Y), 1, painter.arrowColor, false)
}
