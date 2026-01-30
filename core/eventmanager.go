package core

// TODO: separate Observer Pattern into a different package
type Listener struct {
	Event string
}

// any object that implements this method will be a listener
// verified at compile time (TODO: test)
func (lis *Listener) Listen(eventData interface{}) {}

type EventManager struct {
	Listeners []Listener
}

func (e *EventManager) subscribe(event string, l Listener) {
	// stores the event in which the object is subscribing
	l.Event = event
	e.Listeners = append(e.Listeners, l)
}
func (e *EventManager) trigger(event string, eventData interface{}) {
	for _, l := range e.Listeners {
		if l.Event == event {
			l.Listen(eventData)
		}
	}
}
