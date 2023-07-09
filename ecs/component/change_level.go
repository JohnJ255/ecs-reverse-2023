package component

import (
	"github.com/yohamta/donburi"
)

type ChangeLevelData struct {
	Index int
}

var ChangeLevel = donburi.NewComponentType[ChangeLevelData]()
