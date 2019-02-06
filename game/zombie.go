package game

import "math/rand"

type zombieWalkPoint struct {
	x, y Coord
}

type Coord int

// Zombie is a zombie type holding position.
type Zombie struct {
	x, y Coord
	currPos int
	name string
	walkPath []*zombieWalkPoint
}

func newZombie(name string) *Zombie {
	return &Zombie{x: 0, y: 0, name:name}
}

// X returns zombie x pos.
func (z *Zombie) X() Coord {
	return z.x
}

// Y returns zombie y pos.
func (z *Zombie) Y() Coord {
	return z.y
}

// Name returns zombie name.
func (z *Zombie) Name() string {
	return z.name
}

// createWalkPath creates random walking path towards the wall.
func (z *Zombie) createWalkPath(dx, dy int) {
	var path []*zombieWalkPoint
	for y := 0; y < dy; y++ {
		path = append(path, &zombieWalkPoint{
			x: Coord(rangeIn(0, dx - 1)),
			y: Coord(y),
		})
	}
	z.walkPath = path
}

func (z *Zombie) walk() {
	p := z.walkPath[z.currPos]
	z.setPos(p.x, p.y)
	z.currPos++
}

func (z *Zombie) setPos(x, y Coord) {
	z.y = y
	z.x = x
}

func (z *Zombie) atPos(x, y Coord) bool {
	return z.x == x && z.y == y
}

func rangeIn(from, to int) int {
	if from == to {
		panic("rangeIn: from == to")
	}
	if from > to {
		panic("rangeIn: from > to")
	}
	if from < 0 || to < 0 {
		panic("rangeIn: args could not be negative")
	}
	r := from + rand.Intn(to-from)
	return r
}