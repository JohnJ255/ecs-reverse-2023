package system

import (
	"ecs_test_cars/ecs/common"
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/ecs/tags"
	"ecs_test_cars/framework"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type Menu struct {
}

func NewMenu(w donburi.World, name string, pos, size framework.Vec2) *Menu {
	entry := w.Entry(w.Create(component.Menu))
	menu := component.Menu.Get(entry)
	menu.Position = pos
	menu.Size = size
	menu.IsOpened = true
	menu.Caption = name
	menuSystem := &Menu{}
	menu.Elements = []framework.IGUIElement{
		framework.NewButton("New game", func() {
			menu.IsOpened = false
			common.ChangeLevelTo(w, 1)
		}, true),
		framework.NewButton("Restart level", func() {
			menu.IsOpened = false
			level := common.GetCurrentLevel(w)
			common.ChangeLevelTo(w, level.Index)
		}, true),
		framework.NewButton("Next level", func() {
			menu.IsOpened = false
			level := common.GetCurrentLevel(w)
			common.ChangeLevelTo(w, level.Index+1)
		}, true),
		framework.NewHorizontalPanel([]framework.IGUIElement{
			framework.NewButton("-", func() {
				settings := menuSystem.GetSettings(w)
				settings.MasterVolume = framework.Limited(settings.MasterVolume-0.1, 0, 1)
			}, true),
			framework.NewLabel("Sound"),
			framework.NewButton("+", func() {
				settings := menuSystem.GetSettings(w)
				settings.MasterVolume = framework.Limited(settings.MasterVolume+0.1, 0, 1)
			}, true),
		}, []float64{20, 60, 20}),
		framework.NewButton("About", func() {
			menu.IsOpened = false
			common.ChangeLevelTo(w, -1)
		}, true),
	}
	return menuSystem
}

func (m *Menu) Update(ecs *ecs.ECS) {
	menuEntry, ok := donburi.NewQuery(filter.Contains(component.Menu)).First(ecs.World)
	if !ok {
		return
	}
	menu := component.Menu.Get(menuEntry)
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		menu.IsOpened = !menu.IsOpened
	}
	if !menu.IsOpened {
		menuEntry.RemoveComponent(tags.Pause)
		return
	}
	menuEntry.AddComponent(tags.Pause)
	for _, element := range menu.Elements {
		element.Update()
	}
}

func (m *Menu) GetSettings(w donburi.World) *component.SettingsData {
	settingsEntry, ok := donburi.NewQuery(filter.Contains(component.Settings)).First(w)
	if !ok {
		settingsEntry = component.Menu.MustFirst(w)
		settingsEntry.AddComponent(component.Settings)
	}
	return component.Settings.Get(settingsEntry)
}
