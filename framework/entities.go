package framework

import "github.com/hajimehoshi/ebiten/v2"

type Drawing interface {
	GetSprite() *ebiten.Image
	GetTransforms(scale float64) *ebiten.DrawImageOptions
}

type Updating interface {
	Start(f *Framework)
	Update(dt float64)
}

type IContainer interface {
	GetComponents() []IComponent
}

type Entity interface {
	Drawing
	Updating
	IContainer
	GetName() string
	GetModel() Model
}
