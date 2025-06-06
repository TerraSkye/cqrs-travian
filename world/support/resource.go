package support

type Resources [4]int64

func (r Resources) Has(wood, iron, clay, crop int64) bool {
	return r[0] >= wood && r[1] >= iron && r[2] >= clay && r[3] >= crop
}
