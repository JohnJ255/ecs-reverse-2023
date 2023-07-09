package renders

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/framework"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"math"
)

type Sprite struct {
}

func (s *Sprite) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	donburi.NewQuery(filter.Contains(component.Spatial, component.Sprite)).Each(ecs.World, func(e *donburi.Entry) {
		lc, ok := component.CurrentLevel.First(ecs.World)
		if !ok {
			return
		}
		sprite := component.Sprite.Get(e)
		if sprite.Image == nil {
			return
		}
		spatial := component.Spatial.Get(e)
		curLevel := component.CurrentLevel.Get(lc)

		t := s.PivotTransform(spatial.Size, spatial.Pivot, sprite)
		t = s.SpatialTransform(t, spatial.Position, spatial.Rotation)
		t = s.SpatialTransform(t, curLevel.Offset, framework.AngleRight)

		screen.DrawImage(sprite.Image, t)
	})
}

func (s *Sprite) PivotTransform(size framework.Size, pivot framework.VecUV, sprite *component.SpriteData) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Rotate(float64(sprite.Angle))

	spriteSize := sprite.Image.Bounds().Size()
	scaleX := size.Length / (float64(spriteSize.X)*math.Cos(sprite.Angle.F64()) + float64(spriteSize.Y)*math.Sin(sprite.Angle.F64()))
	scaleY := size.Height / (float64(spriteSize.X)*math.Sin(sprite.Angle.F64()) + float64(spriteSize.Y)*math.Cos(sprite.Angle.F64()))
	op.GeoM.Scale(scaleX, scaleY)

	tx := -size.Length * (pivot.U - math.Abs(math.Sin(sprite.Angle.F64())))
	ty := -size.Height * pivot.V
	op.GeoM.Translate(tx, ty)

	return op
}

func (s *Sprite) SpatialTransform(t *ebiten.DrawImageOptions, pos framework.Vec2, rot framework.Radian) *ebiten.DrawImageOptions {
	t.GeoM.Rotate(rot.F64())
	t.GeoM.Translate(pos.X, pos.Y)

	return t
}
