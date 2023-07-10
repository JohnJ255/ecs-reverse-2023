package renders

import (
	"ecs_test_cars/ecs/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"golang.org/x/image/font"
	"image/color"
)

type Menu struct {
	fontGUI font.Face
}

func NewMenu(fontGUI font.Face) *Menu {
	return &Menu{
		fontGUI: fontGUI,
	}
}

func (m *Menu) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	menuEntry, ok := donburi.NewQuery(filter.Contains(component.Menu)).First(ecs.World)
	if !ok {
		return
	}
	menu := component.Menu.Get(menuEntry)
	if !menu.IsOpened {
		return
	}

	x := menu.Position.X
	y := menu.Position.Y
	dy := 40.0
	w := menu.Size.X
	h := menu.Size.Y

	vector.DrawFilledRect(screen, float32(x-50), float32(y-50), float32(w+100), float32(4*(h+dy)), color.NRGBA{100, 100, 100, 150}, false)
	text.Draw(screen, menu.Caption, m.fontGUI, int(x+w/2)-len(menu.Caption)*4, int(y-25), color.White)

	for _, element := range menu.Elements {
		element.Draw(screen, int(x), int(y), int(w), int(h), m.fontGUI)
		y += dy
	}
}
