package game

// Player holds player data.
type Player struct {
	name     string
	notifyCh chan Event
}

const notifyChanEventsBuffer = 100

// newPlayer creates new player instance.
func newPlayer(name string) *Player {
	return &Player{
		name:     name,
		notifyCh: make(chan Event, notifyChanEventsBuffer),
	}
}

// Name returns player name
func (p *Player) Name() string {
	return p.name
}

// Notify returns readonly game notifyCh chan.
func (p *Player) Notify() <-chan Event {
	return p.notifyCh
}

// sendEvent sends game event to underlying
// player notifyCh chan.
func (p *Player) sendEvent(e Event) {
	p.notifyCh <- e
}
