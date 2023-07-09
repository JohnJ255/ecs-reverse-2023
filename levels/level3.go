package levels

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/factory"
	"ecs_test_cars/framework"
	"ecs_test_cars/loader"
	"github.com/yohamta/donburi"
)

type Level3 struct {
	Level1
}

func (l *Level3) Fill(w donburi.World) {
	imgPlayer := framework.InitSprites(framework.Size{10, 10})
	imgPlayer.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Player])
	imgComp := framework.InitSprites(framework.Size{10, 10})
	imgComp.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Computer])
	imgTrailer := framework.InitSprites(framework.Size{10, 10})
	imgTrailer.LoadResources(&loader.ResourceLoader{}, loader.TrailerFileNames[component.TrailerTypeCart])

	car := factory.NewPlayerCar(w, 430, 270, framework.AngleRight, imgPlayer)
	factory.NewTrailerForCar(w, car, imgTrailer)

	factory.NewCar(w, 212, 380, framework.AngleBottom, imgComp)
	factory.NewCar(w, 212, 380, framework.AngleBottom, imgComp)
	factory.NewCar(w, 520, 470, framework.AngleTop+0.07, imgComp)

	l.makeWallsCollisions(w)
	l.makeGoalTrigger(w)
}
