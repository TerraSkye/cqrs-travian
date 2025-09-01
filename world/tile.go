package world

import (
	"fmt"
)

// ================================
// TILE BIT STRUCTURE
// ================================
//
// A Tile is stored as a single uint64 value that packs the following data
// into specific bits:
//
// ┌─────────────┬───────────────────┬───────────────┬─────────────┐
// │ Bit Index   │ Field             │ Size (bits)   │ Notes       │
// ├─────────────┼───────────────────┼───────────────┼─────────────┤
// │ 0 - 2       │ Boost 1           │ 3 bits        │ 0–4 (none + 4 boosts) |
// │ 3 - 5       │ Boost 2           │ 3 bits        │ 0–4         │
// │ 6 - 9       │ Tile Layout ID    │ 4 bits        │ 1–12 valid  │
// │ 10 - 13     │ Display Layout ID │ 4 bits        │ 1–9 valid   │
// │ 14          │ Tile Type         │ 1 bit         │ 0=Village,  │
// │             │                   │               │ 1=Oasis     │
// │ 15+         │ Reserved          │ unused now    │             │
// └─────────────┴───────────────────┴───────────────┴─────────────┘
//
// This compact encoding allows quick storage and comparison of tiles.
// Boosts can be any combination (including duplicates) of the 4 resource boosts.
//
// ================================
// RESOURCE LAYOUTS
// ================================
//
// The Tile Layout ID maps to resource configuration strings:
// For example, layout ID 1 → "3-3-3-9" meaning (wood-clay-iron-crop production)
//
// These are fixed and correspond to map tile resource bonuses.
var tileLayouts = map[uint64]string{
	1:  "3-3-3-9",
	2:  "3-4-5-6",
	3:  "4-4-4-6",
	4:  "4-5-3-6",
	5:  "5-3-4-6",
	6:  "1-1-1-15",
	7:  "4-4-3-7",
	8:  "3-4-4-7",
	9:  "4-3-4-7",
	10: "3-5-4-6",
	11: "4-3-5-6",
	12: "5-4-3-6",
}

//
// ================================
// TYPES AND CONSTANTS
// ================================

// Tile is a packed bitfield storing all tile attributes.
type Tile uint64

// Boost type IDs (fit in 3 bits).
type Boost uint64

const (
	BoostNone Boost = 0
	BoostWood Boost = 1
	BoostClay Boost = 2
	BoostIron Boost = 3
	BoostCrop Boost = 4
)

// TileType indicates whether tile is Village or Oasis.
type TileType uint64

const (
	Village TileType = 0
	Oasis   TileType = 1
)

//
// ================================
// TILE CONSTRUCTOR
// ================================

// NewTile creates a new encoded tile value.
//
// Parameters:
// - tType: 0 = Village, 1 = Oasis
// - layoutID: the tile's resource layout variant (1–12 supported now)
// - displayLayout: the visual arrangement layout ID (1–9 supported now)
// - boost1: first resource boost type
// - boost2: second resource boost type
//
// Notes:
// The order of boosts is not logically important —
// (boost1, boost2) has the same in-game effect as (boost2, boost1) —
// but both are stored explicitly for clarity and potential future mechanics.
//
// Bit packing layout for the encoded Tile (uint64):
//
//	bits  0–2   = boost1 (3 bits)
//	bits  3–5   = boost2 (3 bits)
//	bits  6–9   = layout ID (4 bits)
//	bits 10–13  = display layout ID (4 bits)
//	bit     14  = tile type (1 bit)
//	bits 15+    = reserved (unused)
//
// Returns:
//
//	A Tile value (uint64) containing the packed tile data.
func NewTile(tType TileType, layoutID uint64, displayLayout uint64, boost1, boost2 Boost) Tile {
	val := uint64(tType&1)<<14 |
		(displayLayout&0b1111)<<10 |
		(layoutID&0b1111)<<6 |
		(uint64(boost2)&0b111)<<3 |
		(uint64(boost1) & 0b111)

	return Tile(val)
}

//
// ================================
// TILE ACCESSORS
// ================================

// Type returns the tile's type (Village or Oasis).
func (t Tile) Type() TileType {
	return TileType((t >> 14) & 0b1)
}

// LayoutID returns the numeric resource layout ID (1–12).
func (t Tile) LayoutID() uint64 {
	return (uint64(t) >> 6) & 0b1111
}

// LayoutStr returns the human-readable resource layout string.
func (t Tile) LayoutStr() string {
	return tileLayouts[t.LayoutID()]
}

// DisplayLayout returns the numeric display layout ID (1–9).
func (t Tile) DisplayLayout() uint64 {
	return (uint64(t) >> 10) & 0b1111
}

// Boost1 returns the first resource boost type.
func (t Tile) Boost1() Boost {
	return Boost(t & 0b111)
}

// Boost2 returns the second resource boost type.
func (t Tile) Boost2() Boost {
	return Boost((t >> 3) & 0b111)
}

//
// ================================
// BOOST STRING REPRESENTATION
// ================================

// String returns a human-readable name for the boost type.
func (b Boost) String() string {
	switch b {
	case BoostNone:
		return "None"
	case BoostWood:
		return "Wood"
	case BoostClay:
		return "Clay"
	case BoostIron:
		return "Iron"
	case BoostCrop:
		return "Crop"
	default:
		return "Unknown"
	}
}

//
// ================================
// OPTIONAL STRINGER FOR DEBUGGING
// ================================

// String returns a nicely formatted string representation of the Tile,
// showing all the packed info decoded.
func (t Tile) String() string {
	return fmt.Sprintf(
		"Tile[Type=%v, Layout=%d (%s), DisplayLayout=%d, Boosts=%s & %s]",
		t.Type(), t.LayoutID(), t.LayoutStr(), t.DisplayLayout(), t.Boost1(), t.Boost2(),
	)
}

// NewVillage creates a Village tile with the given layoutID and display layout.
// Villages do not have boosts, so boost1 and boost2 are set to BoostNone.
func NewVillage(layoutID uint64, displayLayout uint64) Tile {
	return NewTile(Village, layoutID, displayLayout, BoostNone, BoostNone)
}

// NewOasis creates an Oasis tile with boosts automatically determined by layoutID.
// Display layout is always fixed to 1 for Oasis tiles.
func NewOasis(layoutID uint64) Tile {
	const oasisDisplayLayout = 1

	// Determine boosts by layoutID according to the mapping
	var boost1, boost2 Boost
	switch layoutID {
	case 1, 2:
		boost1, boost2 = BoostWood, BoostNone
	case 3:
		boost1, boost2 = BoostWood, BoostCrop
	case 4, 5:
		boost1, boost2 = BoostClay, BoostNone
	case 6:
		boost1, boost2 = BoostClay, BoostCrop
	case 7, 8:
		boost1, boost2 = BoostIron, BoostNone
	case 9:
		boost1, boost2 = BoostIron, BoostCrop
	case 10, 11:
		boost1, boost2 = BoostCrop, BoostNone
	case 12:
		boost1, boost2 = BoostCrop, BoostCrop
	default:
		boost1, boost2 = BoostNone, BoostNone
	}

	return NewTile(Oasis, layoutID, oasisDisplayLayout, boost1, boost2)
}
