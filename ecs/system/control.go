package system

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

var Control = func(ecs *ecs.ECS) {
	player, ok := tags.Player.First(ecs.World)
	if !ok {
		return
	}
	car := component.Car.Get(player)

	car.Accelerate = 0.0
	car.WheelRotate = 0.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		car.WheelRotate = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		car.WheelRotate = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		car.Accelerate = 1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		car.Accelerate = -0.3
	}
}
