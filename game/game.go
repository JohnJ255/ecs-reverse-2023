package game

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/renders"
	"ecs_test_cars/ecs/system"
	"ecs_test_cars/framework"
	"ecs_test_cars/loader"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"golang.org/x/image/font"
	"image"
)

const (
	LayerBackground ecs.LayerID = iota
	LayerEntities
	LayerGUI
	LayerMainMenu
	LayersCount
)

type Game struct {
	WindowSize        framework.IntSize
	scale             float64
	f                 *framework.Framework
	fontGUI           font.Face
	fontTTF           *truetype.Font
	SoundMasterVolume float64
	Name              string
	bounds            image.Rectangle
}

func NewGame(ttf *truetype.Font) *Game {
	faceOpt := &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	}

	g := &Game{
		Name: "Reverse to the Garage",
		WindowSize: framework.IntSize{
			Width:  800,
			Height: 600,
		},
		scale:             0.1,
		fontTTF:           ttf,
		fontGUI:           truetype.NewFace(ttf, faceOpt),
		SoundMasterVolume: 0.3,
	}

	return g
}

func (g *Game) Start(f *framework.Framework) {
	f.Audio = framework.NewAudioPlayer(&loader.ResourceLoader{})
	//f.DebugModeEnable()
	f.Audio.SetMasterVolume(g.SoundMasterVolume)
	g.f = f

	g.bounds = image.Rectangle{}

	scenes := system.NewSceneManager()
	audio := system.NewAudioSystem(f.Audio)
	sprite := &renders.Sprite{}
	limiter := system.NewScreenLimiter(g.f)
	camera := system.NewFollowCamera(g.f)
	collision := system.NewCollisionSystem()
	physic := system.NewPhysics()
	gui := renders.NewGUI(g.fontTTF)

	f.Ecs.AddSystem(audio.Update)

	f.Ecs.AddSystem(system.ScoreManager)
	f.Ecs.AddSystem(scenes.Update)
	f.Ecs.AddSystem((&system.CarMoving{}).Update)
	f.Ecs.AddSystem((&system.TrailerMoving{}).Update)
	f.Ecs.AddSystem(system.Control)

	f.Ecs.AddSystem(limiter.Update)
	f.Ecs.AddSystem(collision.Update)
	f.Ecs.AddSystem(physic.Update)

	f.Ecs.AddSystem(camera.Update)

	f.Ecs.AddSystem(system.LevelWin)

	f.Ecs.AddRenderer(LayerBackground, renders.Level)
	f.Ecs.AddRenderer(LayerEntities, sprite.Draw)
	f.Ecs.AddRenderer(LayerGUI, gui.Draw)

	f.Ecs.World.Create(component.CurrentLevel)
}

func (g *Game) Update(dt float64) error {
	g.f.Ecs.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for index := ecs.LayerID(0); index < LayersCount; index++ {
		g.f.Ecs.DrawLayer(index, screen)
	}
}

func (g *Game) GetTitle() string {
	return "ecs_test_cars"
}

func (g *Game) SceneTransform(transforms *ebiten.DrawImageOptions) *ebiten.DrawImageOptions {
	return transforms
}
func (g *Game) IsPaused() bool {
	return false
}
