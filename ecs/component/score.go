package component

import "github.com/yohamta/donburi"

type ScoreData struct {
	Score int
}

var Score = donburi.NewComponentType[ScoreData]()
