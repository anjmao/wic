package game

import (
	"testing"
)

func TestZombieWalk(t *testing.T) {
	points := []*zombieWalkPoint{
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 4, y: 4},
	}

	z := newZombie("z1")
	z.walkPath = points

	for _, p := range points {
		z.walk()
		if z.x != p.x || z.y != p.y {
			t.Fatalf("expected zombie to be at %d:%d, got %d:%d", p.x, p.y, z.x, z.y)
		}
	}
}