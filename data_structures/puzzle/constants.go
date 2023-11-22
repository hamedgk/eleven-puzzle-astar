package puzzle

const (
	Left Direction = iota
	Right
	Up
	Down
	None
)

const (
	Rows  = 3
	Cols  = 4
	Blank = 255
	
	
	MaxUint = ^uint(0) 
	MinUint = 0 
	MaxInt = int(MaxUint >> 1) 
	MinInt = -MaxInt - 1
)
