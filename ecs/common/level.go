package common

import (
	"ecs_test_cars/ecs/component"
	"github.com/yohamta/donburi"
)

func ChangeLevelTo(w donburi.World, index int) {
	levelEntry, ok := component.CurrentLevel.First(w)
	if !ok {
		levelEntry = w.Entry(w.Create(component.CurrentLevel))
	}
	if !levelEntry.HasComponent(component.ChangeLevel) {
		levelEntry.AddComponent(component.ChangeLevel)
	}
	component.ChangeLevel.SetValue(levelEntry, component.ChangeLevelData{
		Index: index,
	})
}

func GetCurrentLevel(w donburi.World) *component.CurrentLevelData {
	levelEntry, ok := component.CurrentLevel.First(w)
	if !ok {
		levelEntry = w.Entry(w.Create(component.CurrentLevel))
	}
	return component.CurrentLevel.Get(levelEntry)
}
