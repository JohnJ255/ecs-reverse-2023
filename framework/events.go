package framework

type Event struct {
	Name string
	Data map[string]interface{}
}

func NewEvent(name string) *Event {
	return &Event{
		Name: name,
		Data: make(map[string]interface{}),
	}
}

type EventSystem struct {
	listeners map[string][]func(event *Event)
}

func NewEventSystem() *EventSystem {
	return &EventSystem{
		listeners: make(map[string][]func(event *Event)),
	}
}

func (es *EventSystem) AddListener(name string, listener func(event *Event)) {
	es.listeners[name] = append(es.listeners[name], listener)
}

func (es *EventSystem) RemoveListeners(name string) {
	delete(es.listeners, name)
}

func (es *EventSystem) Dispatch(event *Event) {
	for _, listener := range es.listeners[event.Name] {
		listener(event)
	}
}
