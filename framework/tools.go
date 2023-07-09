package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"time"
)

type ICollisionFiguresContainer interface {
	GetFigures() []IColliderFigure
}

type ICollisionPainter interface {
	Draw(screen *ebiten.Image, figure IColliderFigure)
}

type IIntersectionPainter interface {
	Draw(screen *ebiten.Image, intersection ContactSet)
}

type DebugTool struct {
	f                  *Framework
	Draws              map[string]func(screen *ebiten.Image)
	CollisionDrawer    ICollisionPainter
	IntersectionDrawer IIntersectionPainter
}

func NewDebugTool(f *Framework, drawer ICollisionPainter, interDrawer IIntersectionPainter) *DebugTool {
	return &DebugTool{
		f:                  f,
		Draws:              make(map[string]func(screen *ebiten.Image)),
		CollisionDrawer:    drawer,
		IntersectionDrawer: interDrawer,
	}
}

func (t *DebugTool) DefaultDrawCollisions(screen *ebiten.Image) {
	for _, ent := range t.f.entities {
		for _, cmp := range ent.GetComponents() {
			switch cmp.(type) {
			case ICollisionFiguresContainer:
				for _, fig := range cmp.(ICollisionFiguresContainer).GetFigures() {
					t.CollisionDrawer.Draw(screen, fig)
				}
			}
		}
	}
}

func (t *DebugTool) DefaultDrawIntersections(csList []ContactSet) func(screen *ebiten.Image) {
	return func(screen *ebiten.Image) {
		for _, cs := range csList {
			t.IntersectionDrawer.Draw(screen, cs)
		}
	}
}

func (t *DebugTool) SetDebugDraw(name string, drawer func(screen *ebiten.Image)) {
	t.Draws[name] = drawer
}

func (t *DebugTool) RemoveDebugDraw(name string) {
	delete(t.Draws, name)
}

func (t *DebugTool) DefaultDrawProjections(index int, pos Vec2, axes []Vec2, pr1 Projection, pr2 Projection) func(screen *ebiten.Image) {
	return func(screen *ebiten.Image) {
		if index < 0 {
			index = (time.Now().Second() / (-index)) % len(axes)
		}
		ax := axes[index]
		var p1, p2 Vec2

		p1 = pos.Sub(ax.Mul(50))
		p2 = p1.Add(ax.Mul(150))
		painter := &DefaultIntersectionPainter{
			arrowColor:  color.Black,
			arrowLength: 10,
			arrowAngle:  Radian(20),
		}
		painter.DrawArrow(screen, p1, p2)
		perp := Vec2{-ax.Y, ax.X}
		length := pos.Length() * math.Sin(float64(pos.ToRadian()-ax.ToRadian()))
		p1 = pos.Sub(perp.Mul(50))
		p2 = p1.Add(perp.Mul(150))
		vector.StrokeLine(screen, float32(p1.X), float32(p1.Y), float32(p2.X), float32(p2.Y), 1, color.NRGBA{200, 200, 255, 155}, false)
		p1 = ax.Mul(pr1.Min).Add(perp.Mul(length))
		p2 = ax.Mul(pr1.Max).Add(perp.Mul(length))
		vector.StrokeLine(screen, float32(p1.X), float32(p1.Y), float32(p2.X), float32(p2.Y), 3, color.NRGBA{40, 40, 255, 100}, false)
		p1 = ax.Mul(pr2.Min).Add(perp.Mul(length))
		p2 = ax.Mul(pr2.Max).Add(perp.Mul(length))
		vector.StrokeLine(screen, float32(p1.X), float32(p1.Y), float32(p2.X), float32(p2.Y), 3, color.NRGBA{255, 40, 40, 100}, false)
	}
}
