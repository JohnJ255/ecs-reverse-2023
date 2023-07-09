package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"image/color"
)

type IGUIElement interface {
	Draw(screen *ebiten.Image, x, y, w, h int)
	Update()
}

type Button struct {
	IsVisible   bool
	IsEnabled   bool
	Text        string
	OnClick     func()
	FillColor   color.Color
	StrokeColor color.Color
	position    Vec2
	size        Size
	fontFace    font.Face
}

func NewButton(text string, onClick func(), fontFace font.Face) *Button {
	return &Button{
		IsVisible:   true,
		IsEnabled:   true,
		Text:        text,
		OnClick:     onClick,
		FillColor:   color.NRGBA{100, 100, 100, 255},
		StrokeColor: color.White,
		fontFace:    fontFace,
	}
}

func (b *Button) Draw(screen *ebiten.Image, x, y, w, h int) {
	b.position = Vec2{float64(x), float64(y)}
	b.size = Size{float64(w), float64(h)}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(w), float32(h), b.FillColor, false)
	vector.StrokeRect(screen, float32(x), float32(y), float32(w), float32(h), 2, b.StrokeColor, false)
	text.Draw(screen, b.Text, b.fontFace, x+w/2-len(b.Text)*4, y+h/2+7, color.NRGBA{230, 240, 250, 255})
}

func (b *Button) Update() {
	if b.IsEnabled && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pos := Vec2{float64(x), float64(y)}
		if pos.InRectWithoutRotation(b.position, b.size) {
			b.OnClick()
		}
	}
}

type Label struct {
	*Button
}

func NewLabel(text string, fontFace font.Face) *Label {
	return &Label{
		&Button{
			IsVisible:   true,
			IsEnabled:   false,
			Text:        text,
			FillColor:   color.NRGBA{100, 100, 100, 255},
			StrokeColor: color.White,
			fontFace:    fontFace,
		},
	}
}

func (l *Label) Draw(screen *ebiten.Image, x, y, w, h int) {
	l.position = Vec2{float64(x), float64(y)}
	l.size = Size{float64(w), float64(h)}
	text.Draw(screen, l.Text, l.fontFace, x+w/2-len(l.Text)*4, y+h/2+7, color.NRGBA{230, 240, 250, 255})
}

func (l *Label) Update() {

}

type Panel struct {
	elements     []IGUIElement
	sizePercents []float64
}

func NewHorizontalPanel(elements []IGUIElement, percents []float64) *Panel {
	return &Panel{
		elements:     elements,
		sizePercents: percents,
	}
}

func (p *Panel) Draw(screen *ebiten.Image, x, y, w, h int) {
	dx := 0
	for i, el := range p.elements {
		percent := 20.0
		if i < len(p.sizePercents) {
			percent = p.sizePercents[i]
		}
		cw := int(float64(w) * percent / 100)
		el.Draw(screen, x+dx, y, cw, h)
		dx += cw
	}
}

func (p *Panel) Update() {
	for _, el := range p.elements {
		el.Update()
	}
}
