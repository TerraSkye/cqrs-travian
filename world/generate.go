package world

import (
	"math/rand"
)

type tileFactory func() Tile

// weightedTile stores the number of times this tile should occur relative to others.
type weightedTile struct {
	weight  int
	factory tileFactory
}

func generateWorld(seed int64, size int) map[Coordinate]Tile {
	seeded := rand.New(rand.NewSource(seed))
	quadrant := size / 2

	tiles := make(map[Coordinate]Tile, (quadrant*2+1)*(quadrant*2+1))

	// Define weights for each tile type
	weightedTiles := []weightedTile{
		// -------------------------------------------------------------------------
		// VILLAGES — Layouts (wood-clay-iron-crop)
		// -------------------------------------------------------------------------
		{10, func() Tile { return NewVillage(1, uint64(seeded.Intn(9)+1)) }},  // 3-3-3-9
		{80, func() Tile { return NewVillage(2, uint64(seeded.Intn(9)+1)) }},  // 3-4-5-6
		{310, func() Tile { return NewVillage(3, uint64(seeded.Intn(9)+1)) }}, // 4-4-4-6
		{80, func() Tile { return NewVillage(4, uint64(seeded.Intn(9)+1)) }},  // 4-5-3-6
		{80, func() Tile { return NewVillage(5, uint64(seeded.Intn(9)+1)) }},  // 5-3-4-6
		{10, func() Tile { return NewVillage(6, uint64(seeded.Intn(9)+1)) }},  // 1-1-1-15
		{30, func() Tile { return NewVillage(7, uint64(seeded.Intn(9)+1)) }},  // 4-4-3-7
		{30, func() Tile { return NewVillage(8, uint64(seeded.Intn(9)+1)) }},  // 3-4-4-7
		{30, func() Tile { return NewVillage(9, uint64(seeded.Intn(9)+1)) }},  // 4-3-4-7
		{80, func() Tile { return NewVillage(10, uint64(seeded.Intn(9)+1)) }}, // 3-5-4-6
		{80, func() Tile { return NewVillage(11, uint64(seeded.Intn(9)+1)) }}, // 4-3-5-6
		{80, func() Tile { return NewVillage(12, uint64(seeded.Intn(9)+1)) }}, // 5-4-3-6

		// -------------------------------------------------------------------------
		// OASES — Boost descriptions
		// -------------------------------------------------------------------------
		{8, func() Tile { return NewOasis(1) }},  // +25% wood
		{8, func() Tile { return NewOasis(2) }},  // +25% wood
		{8, func() Tile { return NewOasis(3) }},  // +25% wood, +25% crop
		{8, func() Tile { return NewOasis(4) }},  // +25% clay
		{8, func() Tile { return NewOasis(5) }},  // +25% clay
		{8, func() Tile { return NewOasis(6) }},  // +25% clay, +25% crop
		{8, func() Tile { return NewOasis(7) }},  // +25% iron
		{8, func() Tile { return NewOasis(8) }},  // +25% iron
		{8, func() Tile { return NewOasis(9) }},  // +25% iron, +25% crop
		{8, func() Tile { return NewOasis(10) }}, // +25% crop
		{8, func() Tile { return NewOasis(11) }}, // +25% crop
		{8, func() Tile { return NewOasis(12) }}, // +25% crop, +25% crop
	}

	// Compute total weight for normalization
	totalWeight := 0
	for _, wt := range weightedTiles {
		totalWeight += wt.weight
	}

	// Generate map tiles
	for x := -quadrant; x <= quadrant; x++ {
		for y := -quadrant; y <= quadrant; y++ {
			coord, err := NewCoordinate(size, x, y)
			if err != nil {
				continue // skip invalid coords
			}

			var tile Tile
			if x == quadrant || x == -quadrant || y == quadrant || y == -quadrant {
				tile = NewVillage(3, uint64(seeded.Intn(9)+1)) // Border always type 3 village
			} else {
				roll := seeded.Intn(totalWeight)
				for _, wt := range weightedTiles {
					if roll < wt.weight {
						tile = wt.factory()
						break
					}
					roll -= wt.weight
				}
			}

			tiles[coord] = tile
		}
	}

	return tiles
}
