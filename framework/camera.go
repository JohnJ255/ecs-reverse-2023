package framework

type ICamera interface {
	Control(obj IGameObject)
	GetPosition() Vec2
}

//todo: make camera as GameEntity

type StaticCamera struct {
	Limit      Size
	Background *Sprite
}

func NewStaticCamera(limit Size, background *Sprite) *StaticCamera {
	return &StaticCamera{
		Limit:      limit,
		Background: background,
	}
}

func (s *StaticCamera) Control(_ IGameObject) {
}

type FollowCamera struct {
	*StaticCamera
	pos Vec2
}

func (f *FollowCamera) GetPosition() Vec2 {
	return f.pos
}

func NewFollowCamera(limit Size, background *Sprite) *FollowCamera {
	return &FollowCamera{
		StaticCamera: &StaticCamera{
			Limit:      limit,
			Background: background,
		},
	}
}

func (f *FollowCamera) Control(obj IGameObject) {
	f.pos = obj.GetPosition().Sub(Vec2{300, 300})
	f.pos.X = Limited(f.pos.X, 0, f.Limit.Length)
	f.pos.Y = Limited(f.pos.Y, 0, f.Limit.Height)
}
