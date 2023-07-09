package framework

import (
	"math"
)

type ITriggerObject interface {
	OnTrigger(entity ICollideOwner, collide *Collide)
}

type ICollideOwner interface {
	IPositioning
	GetScale() Vec2
	GetPivot() VecUV
}

type ISizableColliderOwner interface {
	ISizable
	ICollideOwner
}

type ICollisionComponentOwner interface {
	ISizableColliderOwner
	Updating
}

type IColliderFigure interface {
	Intersect(other IColliderFigure) ContactSet
	Bounds() Bounds
	SetOffset(offset Vec2)
	SetScale(scale Vec2)
	SetRotation(rot Radian)
	GetOwner() ICollideOwner
}

type Collide struct {
	Collision *Collider
	Contacts  []ContactSet
}

func NewCollide(collision *Collider, contacts []ContactSet) *Collide {
	return &Collide{
		Collision: collision,
		Contacts:  contacts,
	}
}

type Collider struct {
	Figures            []IColliderFigure
	BehaviourOnCollide func(collide *Collide)
	entity             ICollideOwner
	f                  *Framework
}

func InitCollider(figure IColliderFigure) *Collider {
	c := &Collider{
		Figures: []IColliderFigure{figure},
	}

	c.BehaviourOnCollide = c.OnCollide

	return c
}

func (c *Collider) GetEntity() ICollideOwner {
	return c.entity
}

func (c *Collider) SetEntity(entity ICollideOwner) {
	c.entity = entity
}

func NewPolygonCollider(points []Vec2, owner ICollideOwner) IColliderFigure {
	return &ColliderShapePolygon{
		points: points,
		owner:  owner,
	}
}

func NewPolygonColliderUV(pointsUV []VecUV, size Size, owner ICollideOwner) IColliderFigure {
	points := make([]Vec2, len(pointsUV))
	for i, p := range pointsUV {
		points[i] = p.ToVec2(size)
	}
	p := &ColliderShapePolygon{
		points: points,
		owner:  owner,
	}
	p.SetOffset(Vec2{}.Sub(owner.GetPivot().ToVec2(size)))

	return p
}

func NewCircleCollider(size Size, owner ICollideOwner) IColliderFigure {
	return &ColliderShapeCircle{
		center: Vec2{size.Length / 2, size.Height / 2},
		radius: math.Min(size.Length, size.Height) / 2,
		owner:  owner,
	}
}

func NewBoxCollider(size Size, owner ICollideOwner) IColliderFigure {
	p1 := VecUV{0, 0}
	p2 := VecUV{1, 0}
	p3 := VecUV{1, 1}
	p4 := VecUV{0, 1}

	return NewPolygonColliderUV([]VecUV{p1, p2, p3, p4}, size, owner)
}

func (c *Collider) AddFigure(f IColliderFigure) {
	c.Figures = append(c.Figures, f)
}

func (c *Collider) GetFigures() []IColliderFigure {
	return c.Figures
}

func (c *Collider) Intersect(collision *Collider) []ContactSet {
	res := make([]ContactSet, 0)
	for _, f := range c.Figures {
		for _, f2 := range collision.Figures {
			contactSet := f.Intersect(f2)
			if contactSet.WasIntersect() {
				res = append(res, contactSet)
			}
		}
	}
	return res
}

func (c *Collider) OnCollide(collide *Collide) {
	if po, ok := c.GetEntity().(IPhysicsObject); ok {
		c.f.Physic.ProcessingCollide(po, collide)
		c.f.Events.Dispatch(&Event{
			Name: "Collider",
			Data: map[string]interface{}{
				"collision": c,
				"collide":   collide,
			},
		})
	}
	if to, ok := collide.Collision.GetEntity().(ITriggerObject); ok {
		to.OnTrigger(c.GetEntity(), collide)
	}
}

func (c *Collider) Start(f *Framework) {
	c.f = f
	f.RegisterCollision(c, c.GetEntity())
}

func (c *Collider) Update(dt float64) {
	collides := make([]*Collide, 0)
	for _, collision := range c.f.GetClosestCollisonsFor(c) {
		cs := c.Intersect(collision)
		if len(cs) > 0 && cs[0].MoveOut != nil {
			collides = append(collides, NewCollide(collision, cs))
		}
	}
	for _, collide := range collides {
		func(collide *Collide) {
			c.f.AddAfterUpdate(func() {
				c.BehaviourOnCollide(collide)
			})
		}(collide)
	}
}

type ContactSet struct {
	Points  []Vec2
	MoveOut *Vec2
	Center  *Vec2
}

func (cs *ContactSet) WasIntersect() bool {
	return cs.Center != nil
}
