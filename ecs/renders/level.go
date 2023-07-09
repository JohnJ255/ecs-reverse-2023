package renders

import (
	"ecs_test_cars/ecs/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var Level = func(ecs *ecs.ECS, screen *ebiten.Image) {
	if entry, ok := donburi.NewQuery(filter.Contains(component.CurrentLevel)).First(ecs.World); ok {
		curLevel := component.CurrentLevel.Get(entry)
		if curLevel.LevelFiller != nil {
			sprite := curLevel.LevelFiller.GetSprite()
			if sprite == nil {
				return
			}
			op := &ebiten.DrawImageOptions{}

			mapSize := curLevel.LevelFiller.GetSize()
			spriteSize := sprite.Bounds().Size()
			scaleX := mapSize.Length / float64(spriteSize.X)
			scaleY := mapSize.Height / float64(spriteSize.Y)
			op.GeoM.Scale(scaleX, scaleY)

			op.GeoM.Translate(curLevel.Offset.X, curLevel.Offset.Y)

			screen.DrawImage(sprite, op)
		}
	}
}
