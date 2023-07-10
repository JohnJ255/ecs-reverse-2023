package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"math"
)

func LevelWin(e *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(e.World); ok {
		return
	}

	donburi.NewQuery(filter.Contains(component.Collision, component.Trigger)).Each(e.World, func(entry *donburi.Entry) {
		t := component.Trigger.Get(entry)
		if t.TriggerType != component.TriggerTypeWin {
			return
		}
		c := component.Collision.Get(entry)

		var sp *component.SpatialData
		tEntry := e.World.Entry(c.WithEntity)
		if tEntry.HasComponent(component.Trailer) && component.Trailer.Get(tEntry).Traktor != nil {
			sp = component.Spatial.Get(tEntry)
			tEntry = e.World.Entry(*component.Trailer.Get(tEntry).Traktor)
		}
		if tEntry.HasComponent(tags.Player) && tEntry.HasComponent(component.Spatial) {
			if sp == nil {
				sp = component.Spatial.Get(tEntry)
			}
			if math.Abs(sp.Rotation.NormalizePi2().F64()) < math.Pi/18 {
				levelEntry := component.CurrentLevel.MustFirst(e.World)
				level := component.CurrentLevel.Get(levelEntry)
				levelEntry.AddComponent(tags.LevelWin)
				levelEntry.AddComponent(component.ChangeLevel)
				component.ChangeLevel.SetValue(levelEntry, component.ChangeLevelData{
					Index: level.Index + 1,
				})
			}
		}
	})

}
