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
	audioEngine     *framework.AudioPlayer
	changeLevelQ    *donburi.Query
	collisionSoundQ *donburi.Query
}

func NewAudioSystem(audioEngine *framework.AudioPlayer) *AudioSystem {
	return &AudioSystem{
		audioEngine:  audioEngine,
		changeLevelQ: donburi.NewQuery(filter.Contains(component.ChangeLevel, tags.LevelWin)),
		collisionSoundQ: donburi.NewQuery(filter.And(
			filter.Contains(component.Physical, component.Spatial, component.Collision),
			filter.Or(filter.Contains(component.Car), filter.Contains(component.Trailer)),
		)),
	}
}

func (s *AudioSystem) Update(ecs *ecs.ECS) {
	if settingsEntry, ok := donburi.NewQuery(filter.Contains(component.Settings)).First(ecs.World); ok {
		settings := component.Settings.Get(settingsEntry)
		s.audioEngine.SetMasterVolume(settings.MasterVolume)
	}

	if level, ok := component.CurrentLevel.First(ecs.World); ok {
		if !level.HasComponent(tags.AudioBackground) {
			level.AddComponent(tags.AudioBackground)
			s.audioEngine.Loop("background")
		}
	}

	s.collisionSoundQ.Each(ecs.World, func(entry *donburi.Entry) {
		s.audioEngine.PlayMany("collide", 200)
	})

	s.changeLevelQ.Each(ecs.World, func(entry *donburi.Entry) {
		s.audioEngine.PlayOnce("win")
	})
}
