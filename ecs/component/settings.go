package component

import "github.com/yohamta/donburi"

type SettingsData struct {
	MasterVolume float64
}

var Settings = donburi.NewComponentType[SettingsData]()
