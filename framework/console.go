package framework

import (
	"ecs_test_cars/helper"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"image/color"
	"strings"
)

type Console struct {
	ToggleKey                  ebiten.Key
	Text                       string
	IsAvailable                bool
	IsOpened                   bool
	inputText                  string
	Padding                    [4]int
	commands                   map[string]func(params ...string) string
	watch                      map[string]func() string
	Background                 color.Color
	Foreground                 color.Color
	IsEnabledDetectMouseClicks bool
	fontTTF                    *truetype.Font
}

func NewConsole(fontTTF *truetype.Font) *Console {
	return &Console{
		ToggleKey:                  ebiten.KeyBackquote,
		Text:                       "",
		IsAvailable:                false,
		IsOpened:                   false,
		inputText:                  "",
		commands:                   make(map[string]func(params ...string) string),
		watch:                      make(map[string]func() string),
		Padding:                    [4]int{10, 10, 10, 10},
		Background:                 color.NRGBA{100, 100, 100, 220},
		Foreground:                 color.NRGBA{220, 220, 220, 255},
		IsEnabledDetectMouseClicks: true,
		fontTTF:                    fontTTF,
	}
}

func (c *Console) Toggle() {
	c.IsOpened = !c.IsOpened
}

func (c *Console) Draw(screen *ebiten.Image, fromX, fromY, toX, toY int) {
	vector.DrawFilledRect(screen, float32(fromX), float32(fromY), float32(toX), float32(toY), c.Background, false)
	faceOpt := &truetype.Options{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	}
	face := truetype.NewFace(c.fontTTF, faceOpt)
	fontHeight := int(face.Metrics().Height) / int(faceOpt.DPI)
	linesCount := strings.Count(c.Text, "\n")
	textHeight := fontHeight * linesCount
	for textHeight > toY-fromY-c.Padding[2]-c.Padding[0]-fontHeight-2*linesCount {
		c.Text = strings.Join(strings.Split(c.Text, "\n")[1:], "\n")
		linesCount = strings.Count(c.Text, "\n")
		textHeight = fontHeight * linesCount
	}
	text.Draw(screen, c.Text, face, fromX+c.Padding[3], fromY+c.Padding[0]+fontHeight, c.Foreground)

	inputFromX := fromX + c.Padding[3]
	inputFromY := toY - c.Padding[2]
	text.Draw(screen, "> "+c.inputText, face, inputFromX, inputFromY, c.Foreground)

	watchFromX := (toX - fromX) * 2 / 3
	watchFromY := fromY + fontHeight + 5 + c.Padding[0]
	text.Draw(screen, "Watch", face, watchFromX, fromY+c.Padding[0]+fontHeight, c.Foreground)
	for lineIndex, wn := range helper.SortedKeys(c.watch) {
		text.Draw(screen, fmt.Sprintf("%s: %s", wn, c.watch[wn]()), face, watchFromX+c.Padding[3], watchFromY+2*c.Padding[0]+(fontHeight+2)*lineIndex, c.Foreground)
	}
	vector.StrokeRect(screen, float32(watchFromX), float32(watchFromY), float32(toX-watchFromX), float32(toY-watchFromY), 2, color.White, false)
}

func (c *Console) Update(f *Framework) {
	k, ok := f.IsPrintableKeyJustPressed()
	if ok && k != c.ToggleKey {
		c.inputText += f.KeyToSymbol(k)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		if len(c.inputText) > 0 {
			c.inputText = c.inputText[:len(c.inputText)-1]
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		cmd := strings.Split(c.inputText, " ")
		c.Println(c.inputText)
		c.makeCommand(cmd[0], cmd[1:]...)
		c.inputText = ""
	}
	if c.IsEnabledDetectMouseClicks && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		c.Println(fmt.Sprintf("x: %d, y: %d", x, y))
	}
}

func (c *Console) makeCommand(name string, params ...string) {
	if do, ok := c.commands[name]; ok {
		c.Println(do(params...))
	} else {
		c.Println("Unknown command: " + name)
	}
}

func (c *Console) SetCommand(name string, do func(params ...string) string) {
	c.commands[name] = do
}

func (c *Console) AddWatch(name string, source func() string) {
	c.watch[name] = source
}

func (c *Console) Println(s string) {
	c.Text += s + "\n"
}
