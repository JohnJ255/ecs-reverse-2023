package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/levels"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

const NewLevelScore = 1000
const CollisionScoreTax = 1

var Scores = func(ecs *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(ecs.World); ok {
		return
	}

	scoreEntry, ok := component.Score.First(ecs.World)
	if !ok {
		scoreEntry, ok = component.CurrentLevel.First(ecs.World)
		if !ok {
			return
		}
		scoreEntry.AddComponent(component.Score)
	}
	score := component.Score.Get(scoreEntry)

	donburi.NewQuery(filter.Contains(component.Physical, component.Spatial, component.Collision)).Each(ecs.World, func(entry *donburi.Entry) {
		isPlayerTrailer := false
		playerCar := tags.Player.MustFirst(ecs.World).Entity()
		if entry.HasComponent(component.Trailer) {
			trailer := component.Trailer.Get(entry)
			isPlayerTrailer = trailer.Traktor != nil && *trailer.Traktor == playerCar
		}
		if playerCar == entry.Entity() || isPlayerTrailer {
			score.Score -= CollisionScoreTax
		}
	})

	entry, ok := tags.LevelStart.First(ecs.World)
	if !ok {
		return
	}

	level := component.CurrentLevel.Get(entry)
	if level.Index == 1 {
		score.Score = 0
	}
	if !entry.HasComponent(tags.LevelWin) && level.Index > 1 {
		return
	}
	entry.RemoveComponent(tags.LevelWin)
	switch level.LevelFiller.(type) {
	case *levels.LevelAbout:
	default:
		score.Score += NewLevelScore
	}
}
