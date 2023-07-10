package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"ecs_test_cars/levels"
	"fmt"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var levelFillers = []component.ILevel{
	nil,
	&levels.Level1{},
	&levels.Level2{},
	&levels.Level3{},
	&levels.Level4{},
	&levels.Level5{},
	&levels.Level6{},
	&levels.LevelAbout{},
}

type LevelsManager struct {
	curLevelQ    *donburi.Query
	changeLevelQ *donburi.Query
}

func NewLevelsManager() *LevelsManager {
	return &LevelsManager{
		curLevelQ:    donburi.NewQuery(filter.Contains(component.CurrentLevel)),
		changeLevelQ: donburi.NewQuery(filter.Contains(component.ChangeLevel)),
	}
}

func (s *LevelsManager) Update(ecs *ecs.ECS) {
	if _, ok := donburi.NewQuery(filter.Contains(component.Menu, tags.Pause)).First(ecs.World); ok {
		return
	}

	if entry, ok := tags.LevelStart.First(ecs.World); ok {
		entry.RemoveComponent(tags.LevelStart)
	}
	if entry, ok := s.changeLevelQ.First(ecs.World); ok {
		newLevel := component.ChangeLevel.Get(entry)
		s.changeLevel(ecs, newLevel)
		entry.RemoveComponent(component.ChangeLevel)
	}
	if entry, ok := s.curLevelQ.First(ecs.World); ok {
		curLevel := component.CurrentLevel.Get(entry)
		if !curLevel.IsFilled {
			if curLevel.LevelFiller == nil {
				entry.AddComponent(component.ChangeLevel)
				component.ChangeLevel.Get(entry).Index = curLevel.Index
				return
			}
			s.removeAllGameObjects(ecs)
			curLevel.LevelFiller.Fill(ecs.World)
			curLevel.IsFilled = true
			entry.AddComponent(tags.LevelStart)
		}
	}
}

func (s *LevelsManager) changeLevel(e *ecs.ECS, newLevel *component.ChangeLevelData) {
	if entry, ok := s.curLevelQ.First(e.World); ok {
		curLevel := component.CurrentLevel.Get(entry)
		curLevel.IsFilled = false
		curLevel.Index = framework.OverLimited(newLevel.Index, 1, len(levelFillers)-1)
		curLevel.LevelFiller = s.chooseFillter(curLevel.Index)
		curLevel.Name = fmt.Sprintf("Level %d", curLevel.Index)
	}
}

func (s *LevelsManager) chooseFillter(index int) component.ILevel {
	if index <= 0 || index > len(levelFillers) {
		index = 1
	}

	return levelFillers[index]
}

func (s *LevelsManager) removeAllGameObjects(e *ecs.ECS) {
	tags.GameElement.Each(e.World, func(entry *donburi.Entry) {
		entry.Remove()
	})
}
