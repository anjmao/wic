package game

import (
	"testing"
	"time"
)

func TestGameZombieMove(t *testing.T) {
	g := createGame()
	go g.Run()

	p := g.JoinPlayer("Donald")
	event := <-p.Notify()
	time.Sleep(1 * time.Millisecond)

	if _, ok := event.(*ZombieMoveEvent); !ok {
		t.Fatalf("expected zombie move event, got %v", event)
	}
}

func TestGameJoinMultiplePlayers(t *testing.T) {
	g := createGame()
	go g.Run()

	g.JoinPlayer("p1")
	g.JoinPlayer("p2")

	expected := 2
	got := len(g.players)
	if expected != got {
		t.Fatalf("expected %d players, got %d", expected, got)
	}
}

func TestGameRemoveRemovePlayersEndsGame(t *testing.T) {
	g := createGame()
	go g.Run()
	g.JoinPlayer("p1")
	g.JoinPlayer("p2")

	g.RemovePlayer("p1")
	g.RemovePlayer("p2")

	expected := 0
	got := len(g.players)
	if expected != got {
		t.Fatalf("expected %d players, got %d", expected, got)
	}
	if !g.ended {
		t.Fatal("expected to end game then all players are removed")
	}
}

func TestGameShootZombie(t *testing.T) {
	g := createGame()
	go g.Run()
	p := g.JoinPlayer("p1")

	g.ShootZombie(p, 0, 0)
	event := <-p.Notify()

	if _, ok := event.(*ShootResultEvent); !ok {
		t.Fatalf("expected shoot event, got %v", event)
	}
	if !g.ended {
		t.Fatal("expected to end game then zombie is killed")
	}
}

func createGame() *Game {
	return NewGame("g1", &Config{Dx: 10, Dy: 30, ZombieMoveInterval: 1 * time.Millisecond})
}
