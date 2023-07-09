package levels

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/factory"
	"ecs_test_cars/framework"
	"ecs_test_cars/loader"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Level1 struct {
}

const WallWidth = 25

func (l *Level1) GetSize() framework.Size {
	return framework.Size{800, 800}
}

func (l Level1) GetSprite() *ebiten.Image {
	sprite := framework.InitSprites(l.GetSize())
	sprite.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[1])
	return sprite.GetSprite()
}

func (l *Level1) Fill(w donburi.World) {
	imgPlayer := framework.InitSprites(framework.Size{10, 10})
	imgPlayer.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Player])
	imgComp := framework.InitSprites(framework.Size{10, 10})
	imgComp.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[framework.Computer])

	factory.NewPlayerCar(w, 330, 265, framework.AngleRight, imgPlayer)

	factory.NewCar(w, 212, 380, framework.AngleBottom, imgComp)

	l.makeWallsCollisions(w)
	l.makeGoalTrigger(w)
}

func (l *Level1) makeWallsCollisions(w donburi.World) {
	factory.NewWall(w, framework.Vec2{114, 200}, framework.Size{130, WallWidth})
	factory.NewWall(w, framework.Vec2{114, 300}, framework.Size{130, WallWidth})
	factory.NewWall(w, framework.Vec2{114, 225}, framework.Size{WallWidth, 75})

	factory.NewWall(w, framework.Vec2{40, 100}, framework.Size{75, 215})
	factory.NewWall(w, framework.Vec2{40, 100}, framework.Size{200, 115})

	factory.NewWall(w, framework.Vec2{40, 70}, framework.Size{720, WallWidth})
	factory.NewWall(w, framework.Vec2{50, 705}, framework.Size{720, WallWidth})
	factory.NewWall(w, framework.Vec2{40, 70}, framework.Size{WallWidth, 705})
	factory.NewWall(w, framework.Vec2{745, 70}, framework.Size{WallWidth, 705})
}

func (l *Level1) makeGoalTrigger(w donburi.World) {
	factory.NewTrigger(w, component.TriggerTypeWin, framework.Vec2{145, 235}, framework.Size{10, 60})
}
