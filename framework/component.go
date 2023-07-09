package framework

type Component struct {
	enabled bool
	owner   Updating
}

func InitComponent() *Component {
	return &Component{
		enabled: true,
	}
}

func (c *Component) IsEnabled() bool {
	return c.enabled
}

func (c *Component) Disable() {
	c.enabled = false
}

func (c *Component) Enable() {
	c.enabled = true
}

func (c *Component) GetOwner() Updating {
	return c.owner
}

func (c *Component) SetOwner(g Updating) {
	c.owner = g
}
