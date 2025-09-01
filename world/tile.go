package world

type Tile interface {
	X() int
	Y() int
	Id() int
	MarshalJSON() ([]byte, error)
}
