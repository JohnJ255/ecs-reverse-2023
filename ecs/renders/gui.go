package renders

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/levels"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"golang.org/x/image/font"
	"image/color"
)

type GUI struct {
	fontGUI font.Face
}

func NewGUI(ttf *truetype.Font) *GUI {
	return &GUI{
		fontGUI: truetype.NewFace(ttf, &truetype.Options{
			Size:    20,
			DPI:     72,
			Hinting: font.HintingFull,
		}),
	}
}

func (g *GUI) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	levelEntry, ok := donburi.NewQuery(filter.Contains(component.CurrentLevel)).First(ecs.World)
	if !ok {
		return
	}
	scoreEntry, ok := component.Score.First(ecs.World)
	if !ok {
		return
	}
	score := component.Score.Get(scoreEntry)
	level := component.CurrentLevel.Get(levelEntry)
	if level != nil {
		switch level.LevelFiller.(type) {
		case *levels.LevelAbout:
		default:
			text.Draw(screen, level.Name, g.fontGUI, 300, 19, color.White)
		}
		text.Draw(screen, fmt.Sprintf("Score: %d", score.Score), g.fontGUI, 400, 19, color.White)
	}
}
