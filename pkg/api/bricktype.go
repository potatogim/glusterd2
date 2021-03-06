package api

// BrickType is the type of Brick
//go:generate jsonenums -type=BrickType
type BrickType uint16

const (
	// Brick represents default type of brick
	Brick BrickType = iota
	// Arbiter represents Arbiter brick type
	Arbiter
)
