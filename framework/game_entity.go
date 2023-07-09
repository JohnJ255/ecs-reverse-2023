package framework

type IPositioning interface {
	GetPosition() Vec2
	GetRotation() Radian
	SetPosition(pos Vec2)
	SetRotation(rot Radian)
}

type ISizable interface {
	GetSize() Size
}

type IGameObject interface {
	IPositioning
	ISizable
}

type IComponent interface {
	GetName() string
	Start(f *Framework)
	Update(dt float64)
	Disable()
	Enable()
	IsEnabled() bool
	GetOwner() Updating
	SetOwner(g Updating)
}

type Model interface {
}

type GameEntity struct {
	Name       string
	Enabled    bool
	Components []IComponent
	Owner      Updating
}

func InitGameEntity(owner Updating) *GameEntity {
	return &GameEntity{
		Components: make([]IComponent, 0),
		Owner:      owner,
	}
}

func (g *GameEntity) Start(f *Framework) {
	for _, c := range g.Components {
		if c.IsEnabled() {
			c.Start(f)
		}
	}
}

func (g *GameEntity) Update(dt float64) {
	for _, c := range g.Components {
		if c.IsEnabled() {
			c.Update(dt)
		}
	}
}

func (g *GameEntity) AddComponent(c IComponent) {
	c.SetOwner(g.Owner)
	g.Components = append(g.Components, c)
}

func (g *GameEntity) GetComponents() []IComponent {
	return g.Components
}

func (g *GameEntity) GetComponent(name string) IComponent {
	for _, c := range g.Components {
		if c.GetName() == name {
			return c
		}
	}
	return nil
}
