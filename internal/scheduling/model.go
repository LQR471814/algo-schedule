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
	PRIORITY_MUST_EXIST
)

type Event struct {
	Id       int
	Name     string
	Priority Priority
	// Duration is in minutes.
	Duration         int
	MinStart, MaxEnd time.Time
}

func (e Event) leeway() int {
	return int(e.MaxEnd.Sub(e.MinStart).Minutes())
}

type TimeBlock struct {
	Start, End time.Time
}

func (b TimeBlock) Duration() time.Duration {
	return b.End.Sub(b.Start)
}

type ScheduleBlock struct {
	TimeBlock
	EventId int
}

// Input assumes that:
//
//   - all the End times of Events are after Now.
//   - all the Deadline times of Tasks are after Now.
type Input struct {
	Now    time.Time
	Events []Event
}

func Schedule(input Input) ([]ScheduleBlock, []error) {
	for i, e := range input.Events {
		if e.MinStart.Before(input.Now) {
			input.Events[i].MinStart = input.Now
		}
	}

	slices.SortFunc(input.Events, func(a, b Event) int {
		if a.Priority < b.Priority {
			return -1
		}
		if a.Priority > b.Priority {
			return 1
		}
		return a.leeway() - b.leeway()
	})

	blocks := make([]ScheduleBlock, 0, len(input.Events))
	for _, e := range input.Events {
		if e.leeway() > 0 {
			continue
		}
		blocks = append(blocks, ScheduleBlock{
			TimeBlock: TimeBlock{
				Start: e.MinStart,
				End:   e.MaxEnd,
			},
			EventId: e.Id,
		})
	}

	cursor := input.Now
	var freeTime []TimeBlock
	for _, e := range input.Events {
		if e.leeway() > 0 {
			continue
		}
		freeMins := int(e.MinStart.Sub(cursor).Minutes())
		if freeMins > 0 {
			freeTime = append(freeTime, TimeBlock{
				Start: cursor,
				End:   e.MinStart,
			})
		}
		cursor = e.MaxEnd
	}

	var errors []error
	for _, e := range input.Events {
		// add inbetween the calendar events (where there is free space)
		for i, free := range freeTime {
			start := free.Start
			if free.Start.Before(e.MinStart) {
				start = e.MinStart
			}
			freeMins := int(free.End.Sub(start))
			if freeMins < e.Duration {
				continue
			}

			end := start.Add(time.Duration(e.Duration) * time.Minute)
			blocks = append(blocks, ScheduleBlock{
				TimeBlock: TimeBlock{
					Start: start,
					End:   end,
				},
				EventId: e.Id,
			})

			if e.MaxEnd.Before(end) {
				errors = append(errors, fmt.Errorf("'%s' was scheduled after its deadline", e.Name))
			}

			freeTimeL := freeTime[:i]
			freeTimeR := freeTime[i+1:]
			remainderL := TimeBlock{
				Start: free.Start,
				End:   start,
			}
			remainderR := TimeBlock{
				Start: end,
				End:   free.End,
			}

			freeTime = freeTimeL
			if remainderL.Duration() > 0 {
				freeTime = append(freeTime, remainderL)
			}
			if remainderR.Duration() > 0 {
				freeTime = append(freeTime, remainderR)
			}
			freeTime = append(freeTime, freeTimeR...)
			break
		}
	}

	slices.SortFunc(blocks, func(a, b ScheduleBlock) int {
		return a.End.Compare(b.End)
	})

	return blocks, errors
}
