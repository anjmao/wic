package game

import (
	"testing"
)

func TestControllerStartGame(t *testing.T) {
	s := NewController()

	g1 := s.StartNewGame(&Config{Dx: 10, Dy: 30})
	g2 := s.StartNewGame(&Config{Dx: 15, Dy: 20})

	if g1 == nil {
		t.Fatal("expected new game g1, got nil")
	}
	if g2 == nil {
		t.Fatal("expected new game g2, got nil")
	}
	expectedGames := 2
	actualGames := len(s.(*controller).games)
	if expectedGames != actualGames {
		t.Fatalf("expected to create %d games, got %d", expectedGames, actualGames)
	}

	// Cleanup.
	g1.end()
	g2.end()
}

func TestControllerDeleteEndedGame(t *testing.T) {
	s := NewController()

	g := s.StartNewGame(&Config{Dx: 10, Dy: 30})
	g.end()

	s.DeleteEndedGame(g.id)

	if len(s.(*controller).games) > 0 {
		t.Fatal("expected to contain no more games")
	}
}