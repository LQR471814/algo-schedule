package scheduling

import (
	"time"
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
	Name string
	// Size is an estimate of how long the task will take in minutes.
	Size int
	// Difficulty measures the decision making involved in this task.
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

func Schedule(input Input) []TimeBlock {
	blocks := make([]TimeBlock, len(input.Events))

	for i, e := range input.Events {
		blocks[i] = TimeBlock{
			Name:  e.Name,
			Start: e.Start,
			End:   e.End,
		}
	}

}
