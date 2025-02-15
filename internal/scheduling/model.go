package scheduling

import "time"

type Size int

const (
	// 15 minutes
	SIZE_SMALL Size = iota
	// 30 minutes
	SIZE_MED
	// 60 minutes
	SIZE_LARGE
)

type Difficulty int

const (
	DIFFICULTY_LOW Difficulty = iota
	DIFFICULTY_MEDIUM
	DIFFICULTY_HIGH
)

type Event struct {
	Name       string
	Start, End time.Time
	Difficulty Difficulty
}

type Quota struct {
	Name       string
	Duration   time.Duration
	Difficulty Difficulty
}

type Task struct {
	Name           string
	Size           Size
	Difficulty     Difficulty
	TimeToDeadline time.Duration
}

type Input struct {
	Tasks  []Task
	Events []Event
	Quotas []Quota
}

type TimeBlock struct {
	Name       string
	Start, End time.Time
}

// func Schedule(input Input) []TimeBlock {
//
// }
