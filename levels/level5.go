package levels

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/factory"
	"ecs_test_cars/framework"
	"ecs_test_cars/loader"
	"github.com/yohamta/donburi"
)

type Level5 struct {
	Level3
}

func (l *Level5) Fill(w donburi.World) {
	imgPlayer := framework.InitSprites(framework.Size{10, 10})
	imgPlayer.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Player])
	imgComp := framework.InitSprites(framework.Size{10, 10})
	imgComp.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Computer])
	imgTrailer := framework.InitSprites(framework.Size{10, 10})
	imgTrailer.LoadResources(&loader.ResourceLoader{}, loader.TrailerFileNames[component.TrailerTypeCart])

	car := factory.NewPlayerCar(w, 430, 450, framework.AngleBottom, imgPlayer)
	factory.NewTrailerForCar(w, car, imgTrailer)

	factory.NewCar(w, 212, 380, framework.AngleBottom, imgComp)
	factory.NewCar(w, 148, 380, framework.AngleBottom, imgComp)
	factory.NewCar(w, 520, 380, framework.AngleBottom, imgComp)
	factory.NewCar(w, 520, 630, framework.AngleTop, imgComp)
	factory.NewCar(w, 650, 640, framework.AngleTop, imgComp)

	l.makeWallsCollisions(w)
	l.makeGoalTrigger(w)
}
