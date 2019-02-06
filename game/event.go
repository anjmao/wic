package game

// Event is a base interface for describing game event.
type Event interface {
	Type() string
}

// ZombieMoveEvent is fired when zombie moves.
type ZombieMoveEvent struct {
	*Zombie
}

func (e ZombieMoveEvent) Type() string { return "ZombieMove" }

// ZombieReachedWallEvent is fired when zombie moves.
type ZombieReachedWallEvent struct {
	*Zombie
}

func (e ZombieReachedWallEvent) Type() string { return "ZombieReachedWall" }

// ShootResultEvent is fired when some player won.
type ShootResultEvent struct {
	Player *Player
	Zombie *Zombie
	Points    int32
}

func (e ShootResultEvent) Type() string { return "ShootResult" }
