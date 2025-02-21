package scheduling

import (
	"fmt"
	"slices"
	"time"
)

type Priority int

const (
	PRIORITY_UNIMPORTANT Priority = iota
	PRIORITY_IMPORTANT
)

type Event struct {
	Name       string
	Start, End time.Time
}

type Task struct {
	Name string
	// Duration is an estimate of how long the task will take in minutes.
	Duration int
	Priority Priority
	Deadline time.Time
}

type TimeBlock struct {
	Start, End time.Time
}

type ScheduleBlock struct {
	TimeBlock
	// Event will be non-nil if the thing scheduled in this time is a Event.
	Event *Event
	// Task will be non-nil if the thing scheduled in this time is a Task.
	Task *Task
}

func (b TimeBlock) Duration() time.Duration {
	return b.End.Sub(b.Start)
}

// TasksIn creates tasks in the intervals provided, this is usually used to convert quotas into tasks.
//
//   - intervals is assumed to be disjoint.
func TasksIn(
	name string,
	duration int,
	priority Priority,
	intervals []TimeBlock,
) []Task {
	tasks := make([]Task, len(intervals))
	for i, intv := range intervals {
		tasks[i] = Task{
			Name:     name,
			Duration: duration,
			Priority: priority,
			Deadline: intv.End,
		}
	}
	return tasks
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

func Schedule(input Input) ([]ScheduleBlock, []error) {
	slices.SortFunc(input.Events, func(a, b Event) int {
		return a.End.Compare(b.End)
	})

	slices.SortFunc(input.Tasks, func(a, b Task) int {
		rA := a.Deadline.Sub(input.Now)
		rB := b.Deadline.Sub(input.Now)
		return int(rA.Minutes() - rB.Minutes())
	})

	blocks := make([]ScheduleBlock, len(input.Events))
	for i, e := range input.Events {
		blocks[i] = ScheduleBlock{
			TimeBlock: TimeBlock{
				Start: e.Start,
				End:   e.End,
			},
			Event: &input.Events[i],
		}
	}

	cursor := input.Now
	var freeTime []TimeBlock
	for _, e := range input.Events {
		freeMins := int(e.Start.Sub(cursor).Minutes())
		if freeMins > 0 {
			freeTime = append(freeTime, TimeBlock{
				Start: cursor,
				End:   e.Start,
			})
		}
		cursor = e.End
	}

	var errors []error
	for ti, t := range input.Tasks {
		// add to the end of all the calendar events
		if len(freeTime) == 0 {
			end := cursor.Add(time.Duration(t.Duration) * time.Minute)
			blocks = append(blocks, ScheduleBlock{
				TimeBlock: TimeBlock{
					Start: cursor,
					End:   end,
				},
				Task: &input.Tasks[ti],
			})
			cursor = end
			continue
		}

		// add inbetween the calendar events (where there is free space)
		for i, free := range freeTime {
			freeMins := int(free.Duration().Minutes())
			if freeMins < t.Duration {
				continue
			}

			if t.Deadline.Before(free.End) {
				errors = append(errors, fmt.Errorf("task '%s' was scheduled after its deadline", t.Name))
			}

			blocks = append(blocks, ScheduleBlock{
				TimeBlock: TimeBlock{
					Start: free.Start,
					End:   free.Start.Add(time.Duration(t.Duration) * time.Minute),
				},
				Task: &input.Tasks[ti],
			})

			remainder := freeMins - t.Duration
			if remainder == 0 {
				freeTime = append(freeTime[:i], freeTime[i+1:]...)
				break
			}
			freeTime[i].Start = free.Start.Add(time.Duration(t.Duration) * time.Minute)
			break
		}
	}

	slices.SortFunc(blocks, func(a, b ScheduleBlock) int {
		return a.End.Compare(b.End)
	})

	return blocks, errors
}
