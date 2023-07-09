package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strings"
)

var printableKeys = map[ebiten.Key][]rune{
	ebiten.KeyA:      {'a'},
	ebiten.KeyB:      {'b'},
	ebiten.KeyC:      {'c'},
	ebiten.KeyD:      {'d'},
	ebiten.KeyE:      {'e'},
	ebiten.KeyF:      {'f'},
	ebiten.KeyG:      {'g'},
	ebiten.KeyH:      {'h'},
	ebiten.KeyI:      {'i'},
	ebiten.KeyJ:      {'j'},
	ebiten.KeyK:      {'k'},
	ebiten.KeyL:      {'l'},
	ebiten.KeyM:      {'m'},
	ebiten.KeyN:      {'n'},
	ebiten.KeyO:      {'o'},
	ebiten.KeyP:      {'p'},
	ebiten.KeyQ:      {'q'},
	ebiten.KeyR:      {'r'},
	ebiten.KeyS:      {'s'},
	ebiten.KeyT:      {'t'},
	ebiten.KeyU:      {'u'},
	ebiten.KeyV:      {'v'},
	ebiten.KeyW:      {'w'},
	ebiten.KeyX:      {'x'},
	ebiten.KeyY:      {'y'},
	ebiten.KeyZ:      {'z'},
	ebiten.KeyDigit1: {'1', '!'},
	ebiten.KeyDigit2: {'2', '@'},
	ebiten.KeyDigit3: {'3', '#'},
	ebiten.KeyDigit4: {'4', '$'},
	ebiten.KeyDigit5: {'5', '%'},
	ebiten.KeyDigit6: {'6', '^'},
	ebiten.KeyDigit7: {'7', '&'},
	ebiten.KeyDigit8: {'8', '*'},
	ebiten.KeyDigit9: {'9', '('},
	ebiten.KeyDigit0: {'0', ')'},
	//ebiten.KeyApostrophe:   {'`', '~'},
	ebiten.KeySpace:        {' '},
	ebiten.KeyEqual:        {'=', '+'},
	ebiten.KeyComma:        {',', '<'},
	ebiten.KeyBackquote:    {'`', '~'},
	ebiten.KeyBackslash:    {'\\', '|'},
	ebiten.KeyBracketLeft:  {'[', '{'},
	ebiten.KeyBracketRight: {']', '}'},
	ebiten.KeySemicolon:    {';', ':'},
	ebiten.KeyQuote:        {'\'', '"'},
	ebiten.KeySlash:        {'/', '?'},
	ebiten.KeyMinus:        {'-', '_'},
}

func (f *Framework) IsAnyKeyJustPressed() bool {
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false
}

func (f *Framework) IsAtLeastOneOfKeyJustPressed(keys []ebiten.Key) (ebiten.Key, bool) {
	for _, k := range keys {
		if inpututil.IsKeyJustPressed(k) {
			return k, true
		}
	}
	return 0, false
}

func (f *Framework) IsPrintableKeyJustPressed() (ebiten.Key, bool) {
	keys := make([]ebiten.Key, len(printableKeys))
	i := 0
	for k := range printableKeys {
		keys[i] = k
		i++
	}
	return f.IsAtLeastOneOfKeyJustPressed(keys)
}

func (f *Framework) KeyToSymbol(k ebiten.Key) string {
	index := 0
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		index = 1
	}
	if s, ok := printableKeys[k]; ok {
		if index < len(s) {
			return string(s[index])
		} else {
			return strings.ToUpper(string(s[0]))
		}
	}
	return ""
}
