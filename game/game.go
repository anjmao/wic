package game

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	DefaultGameConfig = &Config{
		ZombieMoveInterval:2 * time.Second,
		Dx:10,
		Dy:30,
	}
)

// Config holds game configurations.
type Config struct {
	Dx, Dy             int
	ZombieMoveInterval time.Duration
}

// Game is a main struct which holds all game objects.
type Game struct {
	id     string
	config *Config
	ended  bool
	quitCh chan struct{}

	playersMux sync.RWMutex
	players map[string]*Player

	zombieMux sync.RWMutex
	zombie    *Zombie
}

// NewGame returns winter game instance.
func NewGame(id string, config *Config) *Game {
	z := newZombie("night-king")
	z.createWalkPath(config.Dx, config.Dy)

	g := &Game{
		id:      id,
		config:  config,
		players: make(map[string]*Player),
		quitCh:  make(chan struct{}),
		zombie: z,
	}
	return g
}

// JoinPlayer adds new player to the game.
func (g *Game) JoinPlayer(name string) *Player {
	g.playersMux.Lock()
	defer g.playersMux.Unlock()

	p := newPlayer(name)
	g.players[name] = p
	return p
}

// RemovePlayer removes player from the game.
func (g *Game) RemovePlayer(name string) {
	g.playersMux.Lock()
	defer g.playersMux.Unlock()

	delete(g.players, name)

	if len(g.players) == 0 {
		g.end()
	}
}

// ID returns game ID.
func (g *Game) ID() string {
	return g.id
}

// ShootZombie shoots zombie at returns success status if zombie was killed.
func (g *Game) ShootZombie(player *Player, x, y Coord) {
	g.zombieMux.RLock()
	defer g.zombieMux.RUnlock()

	if g.zombie.atPos(x, y) {
		g.sendEvent(&ShootResultEvent{Player: player, Points:1, Zombie:g.zombie})
		g.end()
	} else {
		g.sendEvent(&ShootResultEvent{Player: player, Points:0, Zombie:g.zombie})
	}
}

// Run runs a game loop. Caller is responsible for managing running thread.
func (g *Game) Run() {
	for {
		select {
		case <-g.quitCh:
			logrus.Printf("game %s ended\n", g.id)
			return
		case <-time.Tick(g.config.ZombieMoveInterval):
			g.walkZombie()

			// Print game controller.
			logrus.Println(g)
		}
	}
}

// String implements stringer interface to dump game controller for debugging.
func (g *Game) String() string {
	return fmt.Sprintf(
		"Game: %s, Players count: %d, Zombie %s pos: {x:%d y:%d}",
		g.id,
		len(g.players),
		g.zombie.name,
		g.zombie.x,
		g.zombie.y)
}

func (g *Game) walkZombie() {
	g.zombieMux.Lock()
	defer g.zombieMux.Unlock()

	g.zombie.walk()
	g.sendEvent(&ZombieMoveEvent{g.zombie})

	if int(g.zombie.y) == g.config.Dy - 1 {
		g.sendEvent(&ZombieReachedWallEvent{g.zombie})
		g.end()
	}
}

// end ends game and cleanups main game loop go goroutine.
func (g *Game) end() {
	g.ended = true
	g.quitCh <- struct{}{}
}

// sendEvent sends game notifyCh to all connected players.
func (g *Game) sendEvent(e Event) {
	g.playersMux.RLock()
	players := g.players
	g.playersMux.RUnlock()

	for _, p := range players {
		p.sendEvent(e)
	}
}

var letterRunes = []rune("abcde")

// newGameID generates random 4 chars length GameID for simpler testing.
// On real systems UUID would be a better choose.
func newGameID() string {
	b := make([]rune, 4)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


