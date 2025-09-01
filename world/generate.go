package world

import (
	"math/rand"
)

//generate a world for a given seed and size
func generateWorld(seed int64, size int) map[int]Tile {
	rand.Seed(seed)

	quadrant := size / 2

	var tiles = make(map[int]Tile, 0)

	for x := -quadrant; x <= quadrant; x++ {
		for y := -quadrant; y <= quadrant; y++ {
			coordinate, err := NewCoordinate(size, x, y)
			if err != nil {
				panic(err)
			}
			var atile Tile
			if x == size || x == 0 || y == size || y == 0 {
				atile = NewEmptyTile(coordinate, 3, rand.Intn(9) +1)
			} else {
				random := rand.Intn(1000)

				if random <= 10 {
					atile = NewEmptyTile(coordinate, 1, rand.Intn(9) +1)
				} else if random <= 90 {
					atile = NewEmptyTile(coordinate, 2, rand.Intn(9) +1)
				} else if random <= 400 {
					atile = NewEmptyTile(coordinate, 3, rand.Intn(9) +1)
				} else if random <= 480 {
					atile = NewEmptyTile(coordinate, 4, rand.Intn(9) +1)
				} else if random <= 560 {
					atile = NewEmptyTile(coordinate, 5, rand.Intn(9) +1)
				} else if random <= 570 {
					atile = NewEmptyTile(coordinate, 6, rand.Intn(9) +1)
				} else if random <= 600 {
					atile = NewEmptyTile(coordinate, 7, rand.Intn(9) +1)
				} else if random <= 630 {
					atile = NewEmptyTile(coordinate, 8, rand.Intn(9) +1)
				} else if random <= 660 {
					atile = NewEmptyTile(coordinate, 9, rand.Intn(9) +1)
				} else if random <= 740 {
					atile = NewEmptyTile(coordinate, 10, rand.Intn(9) +1)
				} else if random <= 820 {
					atile = NewEmptyTile(coordinate, 11, rand.Intn(9) +1)
				} else if random <= 900 {
					atile = NewEmptyTile(coordinate, 12, rand.Intn(9) +1)
				} else if random <= 908 {
					atile = NewOasisTile(coordinate, 1, rand.Intn(9) +1)
				} else if random <= 916 {
					atile = NewOasisTile(coordinate, 2, rand.Intn(9) +1)
				} else if random <= 924 {
					atile = NewOasisTile(coordinate, 3, rand.Intn(9) +1)
				} else if random <= 932 {
					atile = NewOasisTile(coordinate, 4, rand.Intn(9) +1)
				} else if random <= 940 {
					atile = NewOasisTile(coordinate, 5, rand.Intn(9) +1)
				} else if random <= 948 {
					atile = NewOasisTile(coordinate, 6, rand.Intn(9) +1)
				} else if random <= 956 {
					atile = NewOasisTile(coordinate, 7, rand.Intn(9) +1)
				} else if random <= 964 {
					atile = NewOasisTile(coordinate, 8, rand.Intn(9) +1)
				} else if random <= 972 {
					atile = NewOasisTile(coordinate, 9, rand.Intn(9) +1)
				} else if random <= 980 {
					atile = NewOasisTile(coordinate, 10, rand.Intn(9) +1)
				} else if random <= 988 {
					atile = NewOasisTile(coordinate, 11, rand.Intn(9) +1)
				} else if random <= 1000 {
					atile = NewOasisTile(coordinate, 12, rand.Intn(9) +1)
				}

			}
			tiles[coordinate.Id()] = atile
		}
	}

	return tiles
}
