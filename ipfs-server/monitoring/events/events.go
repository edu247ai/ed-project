package events

type Event struct {
	Type string
	Data map[string]interface{}
}

type Bus struct {
	subscribers []chan Event
}

func NewBus() *Bus {
	return &Bus{subscribers: []chan Event{}}
}

func (b *Bus) Subscribe() chan Event {
	ch := make(chan Event, 100)
	b.subscribers = append(b.subscribers, ch)
	return ch
}

func (b *Bus) Publish(e Event) {
	for _, ch := range b.subscribers {
		ch <- e
	}
}
