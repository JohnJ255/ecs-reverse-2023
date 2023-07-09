package component

import "github.com/yohamta/donburi"

type TriggerType int

const (
	TriggerTypeWin TriggerType = iota
)

type TriggerData struct {
	TriggerType TriggerType
}

var Trigger = donburi.NewComponentType[TriggerData]()
