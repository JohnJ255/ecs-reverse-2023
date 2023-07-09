package framework

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"time"
)

const sampleRate = 44100

type IAudioResourceLoader interface {
	GetSound(filename string, audioContext *audio.Context) *audio.Player
	GetSoundList() []string
}

type AudioPlayer struct {
	audioContext *audio.Context
	sounds       map[string]*audio.Player
	loop         []*audio.Player
	loader       IAudioResourceLoader
	freq         map[string]int64
}

func NewAudioPlayer(loader IAudioResourceLoader) *AudioPlayer {
	a := &AudioPlayer{
		audioContext: audio.NewContext(sampleRate),
		sounds:       make(map[string]*audio.Player),
		loop:         make([]*audio.Player, 0),
		loader:       loader,
		freq:         make(map[string]int64),
	}
	a.preload()
	return a
}

func (a *AudioPlayer) Play(name string) {
	if _, ok := a.sounds[name]; !ok {
		a.sounds[name] = a.loader.GetSound(name, a.audioContext)
	}
	a.sounds[name].Rewind()
	a.sounds[name].Play()
}

func (a *AudioPlayer) PlayMany(name string, freqMilliseconds int64) {
	if _, ok := a.freq[name]; !ok {
		a.freq[name] = time.Now().UnixMilli()
	}
	volume := 1.0
	if _, ok := a.sounds[name]; ok {
		volume = a.sounds[name].Volume()
	}
	if time.Now().UnixMilli()-a.freq[name] >= freqMilliseconds {
		a.freq[name] = time.Now().UnixMilli()
		sound := a.loader.GetSound(name, a.audioContext)
		sound.SetVolume(volume)
		sound.Play()
	}
}

func (a *AudioPlayer) PlayOnce(name string) {
	if _, ok := a.sounds[name]; !ok {
		a.sounds[name] = a.loader.GetSound(name, a.audioContext)
	}
	if !a.sounds[name].IsPlaying() {
		a.sounds[name].Rewind()
		a.sounds[name].Play()
	}
}

func (a *AudioPlayer) StopAll() {
	a.loop = make([]*audio.Player, 0)
	for _, player := range a.sounds {
		if player.IsPlaying() {
			player.Pause()
			player.Rewind()
		}
	}
}

func (a *AudioPlayer) Loop(name string) {
	a.Play(name)
	if _, ok := a.sounds[name]; ok {
		a.loop = append(a.loop, a.sounds[name])
	}
}

func (a *AudioPlayer) Update() {
	for _, player := range a.loop {
		if !player.IsPlaying() {
			player.Rewind()
			player.Play()
		}
	}
}

func (a *AudioPlayer) SetVolume(name string, volume float64) {
	if _, ok := a.sounds[name]; !ok {
		a.sounds[name] = a.loader.GetSound(name, a.audioContext)
	}
	a.sounds[name].SetVolume(volume)
}

func (a *AudioPlayer) preload() {
	for _, name := range a.loader.GetSoundList() {
		a.sounds[name] = a.loader.GetSound(name, a.audioContext)
	}
}

func (a *AudioPlayer) SetMasterVolume(volume float64) {
	for _, player := range a.sounds {
		player.SetVolume(volume)
	}
}
