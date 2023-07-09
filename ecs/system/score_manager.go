package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/levels"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

const NewLevelScore = 1000

var ScoreManager = func(ecs *ecs.ECS) {
	scoreEntry, ok := component.Score.First(ecs.World)
	if !ok {
		scoreEntry, ok = component.CurrentLevel.First(ecs.World)
		if !ok {
			return
		}
		scoreEntry.AddComponent(component.Score)
	}
	score := component.Score.Get(scoreEntry)

	tags.LevelStart.Each(ecs.World, func(entry *donburi.Entry) {
		level := component.CurrentLevel.Get(entry)
		if level.Index == 1 {
			score.Score = 0
		}
		switch level.LevelFiller.(type) {
		case *levels.LevelAbout:
		default:
			score.Score += NewLevelScore
		}
	})
}
