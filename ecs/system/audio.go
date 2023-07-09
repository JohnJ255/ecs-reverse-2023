package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type AudioSystem struct {
	audioEngine  *framework.AudioPlayer
	changeLevelQ *donburi.Query
}

func NewAudioSystem(audioEngine *framework.AudioPlayer) *AudioSystem {
	return &AudioSystem{
		audioEngine:  audioEngine,
		changeLevelQ: donburi.NewQuery(filter.Contains(component.ChangeLevel, tags.LevelFinish)),
	}
}

func (s *AudioSystem) Update(ecs *ecs.ECS) {
	if level, ok := component.CurrentLevel.First(ecs.World); ok {
		if !level.HasComponent(tags.AudioBackground) {
			level.AddComponent(tags.AudioBackground)
			s.audioEngine.Loop("background")
		}
	}
	s.changeLevelQ.Each(ecs.World, func(entry *donburi.Entry) {
		s.audioEngine.PlayOnce("win")
	})
}
