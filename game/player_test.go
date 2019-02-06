package game

import "testing"

func TestPlayerEventsChan(t *testing.T) {
	p := newPlayer("Donald")

	wonEvent := &ShootResultEvent{Player:p}
	p.sendEvent(wonEvent)
	event := <-p.Notify()

	if event != wonEvent {
		t.Fatalf("expected wonEvent, got %v", event)
	}
}
