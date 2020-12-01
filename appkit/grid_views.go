package appkit

type GridRowAlignment int

const (
	GridRowAlignmentInherited GridRowAlignment = iota
	GridRowAlignmentNone
	GridRowAlignmentFirstBaseline
	GridRowAlignmentLastBaseline
)

type GridCellPlacement int

const (
	GridCellPlacementInherited = iota
	GridCellPlacementNone
	GridCellPlacementLeading

	GridCellPlacementTrailing

	GridCellPlacementCenter
	GridCellPlacementFill
	GridCellPlacementTop    = GridCellPlacementLeading
	GridCellPlacementBottom = GridCellPlacementTrailing
)
