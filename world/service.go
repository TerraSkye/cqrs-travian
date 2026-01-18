package world

import (
	"errors"
)

var (
	TileOutOfBounds = errors.New("world is out of bound")
)

type World interface {
	ViewCluster(center Coordinate, zoom int) (map[int]map[int]Tile, error)
	ViewTile(id Coordinate) (Tile, error)
	GetCoordinateById(id int) (Coordinate, error)
	GetCoordinate(x int, y int) (Coordinate, error)
	Size() int
}

type world struct {
	seed  int64
	size  int
	world map[Coordinate]Tile // this is the empty world for the given seed.
}

// generate a new world for a given seed with a size.
func NewWorld(seed int64, size int) World {
	return &world{
		seed:  seed,
		size:  size,
		world: generateWorld(seed, size),
	}
}

func (t *world) Size() int {
	return t.size
}

func (t *world) ViewTile(coordinate Coordinate) (Tile, error) {
	return t.world[coordinate], nil
}

func (t *world) GetCoordinate(x int, y int) (Coordinate, error) {
	return NewCoordinate(t.size, x, y)
}

func (t *world) GetCoordinateById(id int) (Coordinate, error) {
	if id > (t.size*t.size - 1) {
		return Coordinate{}, TileOutOfBounds
	}

	equator := t.size / 2

	x := abs(t.size, id/t.size+equator+1)
	y := abs(t.size, id%t.size-equator)

	return NewCoordinate(t.size, x, y)

}

func (t *world) ViewCluster(centeredAt Coordinate, zoom int) (map[int]map[int]Tile, error) {
	tiles := make(map[int]map[int]Tile, 0)

	for x := centeredAt.x - (zoom / 2); x < centeredAt.x+(zoom/2)+1; x++ {
		tiles[abs(t.size, x)] = make(map[int]Tile, 0)
		for y := centeredAt.y - (zoom / 2); y < centeredAt.y+(zoom/2)+1; y++ {
			pos, _ := NewCoordinate(t.size, x, y)
			tiles[pos.X()][pos.Y()] = t.world[pos]
		}
	}
	return tiles, nil
}
