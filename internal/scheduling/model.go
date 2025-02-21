package scheduling

import (
	"slices"
	"time"
)

type Priority int

const (
	PRIORITY_UNIMPORTANT Priority = iota
	PRIORITY_IMPORTANT
)

type Event struct {
	Id         uint64
	Name       string
	Start, End time.Time
}

type Task struct {
	Id   uint64
	Name string
	// Duration is an estimate of how long the task will take in minutes.
	Duration int
	Priority Priority
	Deadline time.Time
}

// Input assumes that:
//
//   - all the End times of Events are after Now.
//   - all the Deadline times of Tasks are after Now.
type Input struct {
	Now    time.Time
	Events []Event
	Tasks  []Task
}

type TimeBlock struct {
	Name       string
	Start, End time.Time
}

func Schedule(input Input) []TimeBlock {
	var blocks []TimeBlock

	slices.SortFunc(input.Tasks, func(a, b Task) int {
		rA := a.Deadline.Sub(input.Now)
		rB := b.Deadline.Sub(input.Now)
		return int(rA.Seconds() - rB.Seconds())
	})

	var prevEnd time.Time
	for i, e := range input.Events {
		blocks[i] = TimeBlock{
			Name:  e.Name,
			Start: e.Start,
			End:   e.End,
		}
		prevEnd = e.End
	}

}
