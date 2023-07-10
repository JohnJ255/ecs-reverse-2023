package system

import (
	"ecs_test_cars/ecs/common"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi/ecs"
)

func HotKeys(ecs *ecs.ECS) {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		level := common.GetCurrentLevel(ecs.World)
		common.ChangeLevelTo(ecs.World, level.Index+1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		level := common.GetCurrentLevel(ecs.World)
		common.ChangeLevelTo(ecs.World, level.Index)
	}
}
