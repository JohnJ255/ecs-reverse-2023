package levels

import (
	"ecs_test_cars/ecs/factory"
	"ecs_test_cars/framework"
	"ecs_test_cars/loader"
	"github.com/yohamta/donburi"
)

type Level2 struct {
	Level1
}

func (l *Level2) Fill(w donburi.World) {
	imgPlayer := framework.InitSprites(framework.Size{10, 10})
	imgPlayer.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Player])
	imgComp := framework.InitSprites(framework.Size{10, 10})
	imgComp.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Computer])

	factory.NewPlayerCar(w, 330, 365, framework.AngleBottom, imgPlayer)

	factory.NewCar(w, 212, 380, framework.AngleBottom, imgComp)
	factory.NewCar(w, 520, 380, framework.AngleBottom, imgComp)

	l.makeWallsCollisions(w)
	l.makeGoalTrigger(w)
}
