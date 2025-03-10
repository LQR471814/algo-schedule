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

type Reservable struct {
	Name     string
	Priority Priority
	// Duration is in minutes.
	Duration         int
	MinStart, MaxEnd time.Time
}

func (e Reservable) leeway() int {
	return int(e.MaxEnd.Sub(e.MinStart).Minutes()) - e.Duration
}

type TimeBlock struct {
	Start, End time.Time
}

func (b TimeBlock) Duration() time.Duration {
	return b.End.Sub(b.Start)
}

type ScheduleBlock struct {
	TimeBlock
	Reservable *Reservable
}

// Input assumes that:
//
//   - all the End times of Events are after Now.
//   - all the Deadline times of Tasks are after Now.
type Input struct {
	Now         time.Time
	Reservables []Reservable
}

func Schedule(input Input) ([]ScheduleBlock, []error) {
	slices.SortFunc(input.Reservables, func(a, b Reservable) int {
		// schedule highest priority first
		if a.Priority < b.Priority {
			return 1
		}
		if a.Priority > b.Priority {
			return -1
		}
		return a.leeway() - b.leeway()
	})

	blocks := make([]ScheduleBlock, 0, len(input.Reservables))
	for i, r := range input.Reservables {
		if r.leeway() != 0 {
			continue
		}
		blocks = append(blocks, ScheduleBlock{
			TimeBlock: TimeBlock{
				Start: r.MinStart,
				End:   r.MaxEnd,
			},
			Reservable: &input.Reservables[i],
		})
	}

	slices.SortFunc(blocks, func(a, b ScheduleBlock) int {
		return a.End.Compare(b.End)
	})
	cursor := input.Now
	var freeTime []TimeBlock
	for _, b := range blocks {
		freeMins := int(b.End.Sub(cursor).Minutes())
		if freeMins <= 0 {
			continue
		}
		freeTime = append(freeTime, TimeBlock{
			Start: cursor,
			End:   b.Start,
		})
		cursor = b.End
	}
	freeTime = append(freeTime, TimeBlock{
		Start: cursor,
		// this is the maximum unix time
		End: time.Unix(1<<63-62135596801, 999999999),
	})
	for _, t := range freeTime {
		fmt.Println(t.Start.Format(time.DateTime), t.End.Format(time.DateTime))
	}

	var errors []error
	for ri, r := range input.Reservables {
		if r.leeway() == 0 {
			continue
		}

		// add inbetween the calendar events (where there is free space)
		for i, free := range freeTime {
			start := free.Start
			if free.Start.Before(r.MinStart) {
				start = r.MinStart
			}
			freeMins := free.End.Sub(start)
			if freeMins < time.Duration(r.Duration)*time.Minute {
				continue
			}

			end := start.Add(time.Duration(r.Duration) * time.Minute)
			fmt.Println("schedule", r.Name, r.Duration, "|", start.Format(time.DateTime), "|", end.Format(time.DateTime))
			blocks = append(blocks, ScheduleBlock{
				TimeBlock: TimeBlock{
					Start: start,
					End:   end,
				},
				Reservable: &input.Reservables[ri],
			})

			if r.MaxEnd.Before(end) {
				errors = append(errors, fmt.Errorf("'%s' was scheduled after its deadline", r.Name))
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

func Event(name string, start, end time.Time) Reservable {
	return Reservable{
		Name:     fmt.Sprintf("Event: %s", name),
		Priority: PRIORITY_MUST_EXIST,
		MinStart: start,
		MaxEnd:   end,
		Duration: int(end.Sub(start).Minutes()),
	}
}

func Task(name string, priority Priority, duration int, deadline time.Time) Reservable {
	return Reservable{
		Name:     fmt.Sprintf("Task: %s", name),
		Priority: priority,
		MinStart: time.Time{},
		MaxEnd:   deadline,
		Duration: duration,
	}
}
