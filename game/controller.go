package game

import (
	"sync"
)

// Controller holds all running games state.
// It is responsible for creating new games.
type Controller interface {
	StartNewGame(config *Config) *Game
	GetGame(id string) *Game
	GetGameIDs() []string
	DeleteEndedGame(id string)
}

type controller struct {
	gamesMux sync.RWMutex
	games    map[string]*Game
}

// NewController creates controller instance.
func NewController() Controller {
	s := &controller{
		games: make(map[string]*Game),
	}
	return s
}

// StartNewGame adds new game to the controller and starts
// it's main loop.
func (s *controller) StartNewGame(config *Config) *Game {
	s.gamesMux.Lock()
	defer s.gamesMux.Unlock()

	id := newGameID()
	g := NewGame(id, config)
	s.games[g.id] = g

	// Start main game loop in separate go routine.
	go g.Run()

	return g
}

// GetGame returns running game from the controller.
func (s *controller) GetGame(id string) *Game {
	s.gamesMux.RLock()
	defer s.gamesMux.RUnlock()

	return s.games[id]
}

// GetGameIDs reads all running games ids.
func (s *controller) GetGameIDs() []string {
	s.gamesMux.RLock()
	defer s.gamesMux.RUnlock()

	var res []string
	for _, g := range s.games {
		res = append(res, g.ID())
	}
	return res
}

// DeleteEndedGame removes game if it's ended.
func (s *controller) DeleteEndedGame(id string) {
	s.gamesMux.Lock()
	defer s.gamesMux.Unlock()

	if g, ok := s.games[id]; ok && g.ended {
		delete(s.games, id)
	}
}
