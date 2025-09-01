package world

import "testing"

// TestNewTile_EncodingDecoding ensures NewTile packs and unpacks fields correctly.
func TestNewTile_EncodingDecoding(t *testing.T) {
	tests := []struct {
		name          string
		tType         TileType
		layoutID      uint64
		displayLayout uint64
		boost1        Boost
		boost2        Boost
	}{
		{"Village basic", Village, 1, 1, BoostWood, BoostClay},
		{"Oasis crop/iron", Oasis, 12, 9, BoostCrop, BoostIron},
		{"Village double crop", Village, 6, 5, BoostCrop, BoostCrop},
		{"Oasis none/wood", Oasis, 3, 4, BoostNone, BoostWood},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tile := NewTile(tt.tType, tt.layoutID, tt.displayLayout, tt.boost1, tt.boost2)

			if tile.Type() != tt.tType {
				t.Errorf("Type() = %v, want %v", tile.Type(), tt.tType)
			}
			if tile.LayoutID() != tt.layoutID {
				t.Errorf("LayoutID() = %v, want %v", tile.LayoutID(), tt.layoutID)
			}
			if tile.DisplayLayout() != tt.displayLayout {
				t.Errorf("DisplayLayout() = %v, want %v", tile.DisplayLayout(), tt.displayLayout)
			}
			if tile.Boost1() != tt.boost1 {
				t.Errorf("Boost1() = %v, want %v", tile.Boost1(), tt.boost1)
			}
			if tile.Boost2() != tt.boost2 {
				t.Errorf("Boost2() = %v, want %v", tile.Boost2(), tt.boost2)
			}

			expectedLayout := tileLayouts[tt.layoutID]
			if tile.LayoutStr() != expectedLayout {
				t.Errorf("LayoutStr() = %v, want %v", tile.LayoutStr(), expectedLayout)
			}
		})
	}
}

// TestBoostString ensures Boost.String returns correct names.
func TestBoostString(t *testing.T) {
	cases := map[Boost]string{
		BoostNone: "None",
		BoostWood: "Wood",
		BoostClay: "Clay",
		BoostIron: "Iron",
		BoostCrop: "Crop",
	}

	for boost, expected := range cases {
		if got := boost.String(); got != expected {
			t.Errorf("Boost(%v).String() = %q, want %q", boost, got, expected)
		}
	}
}

// TestNewTile_AllCombinations runs through all legal combinations to verify consistency.
func TestNewTile_AllCombinations(t *testing.T) {
	for _, tType := range []TileType{Village, Oasis} {
		for layoutID := range tileLayouts {
			for displayLayout := uint64(1); displayLayout <= 9; displayLayout++ {
				for boost1 := BoostNone; boost1 <= BoostCrop; boost1++ {
					for boost2 := BoostNone; boost2 <= BoostCrop; boost2++ {
						tile := NewTile(tType, uint64(layoutID), displayLayout, boost1, boost2)

						if tile.Type() != tType {
							t.Fatalf("Mismatch: got Type=%v, want %v", tile.Type(), tType)
						}
						if tile.LayoutID() != uint64(layoutID) {
							t.Fatalf("Mismatch: got LayoutID=%v, want %v", tile.LayoutID(), layoutID)
						}
						if tile.DisplayLayout() != displayLayout {
							t.Fatalf("Mismatch: got DisplayLayout=%v, want %v", tile.DisplayLayout(), displayLayout)
						}
						if tile.Boost1() != boost1 {
							t.Fatalf("Mismatch: got Boost1=%v, want %v", tile.Boost1(), boost1)
						}
						if tile.Boost2() != boost2 {
							t.Fatalf("Mismatch: got Boost2=%v, want %v", tile.Boost2(), boost2)
						}
					}
				}
			}
		}
	}
}
