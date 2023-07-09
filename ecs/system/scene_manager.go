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

var levelFillters = []component.ILevel{
	nil,
	&levels.Level1{},
	&levels.Level2{},
	&levels.Level3{},
	&levels.Level4{},
	&levels.Level5{},
	&levels.Level6{},
	&levels.LevelAbout{},
}

type SceneManager struct {
	curLevelQ    *donburi.Query
	changeLevelQ *donburi.Query
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		curLevelQ:    donburi.NewQuery(filter.Contains(component.CurrentLevel)),
		changeLevelQ: donburi.NewQuery(filter.Contains(component.ChangeLevel)),
	}
}

func (s *SceneManager) Update(ecs *ecs.ECS) {
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

func (s *SceneManager) changeLevel(e *ecs.ECS, newLevel *component.ChangeLevelData) {
	if entry, ok := s.curLevelQ.First(e.World); ok {
		curLevel := component.CurrentLevel.Get(entry)
		curLevel.IsFilled = false
		curLevel.Index = framework.Limited(newLevel.Index, 1, len(levelFillters)-1)
		curLevel.LevelFiller = s.chooseFillter(curLevel.Index)
		curLevel.Name = fmt.Sprintf("Level %d", curLevel.Index)

		//q := donburi.NewQuery(filter.Contains(component.GameElementTag))
		//q.Each(e.World, func(entry *donburi.Entry) {
		//	e.World.Remove(entry.Entity())
		//})
	}
}

func (s *SceneManager) chooseFillter(index int) component.ILevel {
	if index <= 0 || index > len(levelFillters) {
		index = 1
	}

	return levelFillters[index]
}

func (s *SceneManager) removeAllGameObjects(e *ecs.ECS) {
	tags.GameElement.Each(e.World, func(entry *donburi.Entry) {
		entry.Remove()
	})
}
