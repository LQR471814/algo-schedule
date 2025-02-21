package scheduling

import (
	"algo-schedule/internal/templates"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

func aDay(now time.Time) ([]Event, []Task) {
	return []Event{
			{
				Name:  "AP Statistics",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, time.Local),
			},
			{
				Name:  "AP Microeconomics",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 9, 25, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, time.Local),
			},
			{
				Name:  "Multi-Variable Calculus (H)",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 10, 55, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, time.Local),
			},
			{
				Name:  "Lunch",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 12, 30, 0, 0, time.Local),
			},
			{
				Name:  "World Religions",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 14, 0, 0, 0, time.Local),
			},
		}, []Task{
			{
				Name:     "AP Stats Webassign",
				Duration: 60,
				Priority: PRIORITY_IMPORTANT,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+2, 8, 0, 0, 0, time.Local),
			},
			{
				Name:     "Multi Webassign",
				Duration: 90,
				Priority: PRIORITY_IMPORTANT,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+2, 8, 0, 0, 0, time.Local),
			},
			{
				Name:     "Econ Assignment: Part 1 / 2",
				Duration: 90,
				Priority: PRIORITY_IMPORTANT,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+7, 8, 0, 0, 0, time.Local),
			},
			{
				Name:     "Econ Assignment: Part 2 / 2",
				Duration: 90,
				Priority: PRIORITY_IMPORTANT,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+7, 8, 0, 0, 0, time.Local),
			},
		}
}

func bDay(now time.Time) ([]Event, []Task) {
	return []Event{
			{
				Name:  "AP Physics",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, time.Local),
			},
			{
				Name:  "Philosophy in Literature (H)",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 9, 25, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, time.Local),
			},
			{
				Name:  "Data Structures and Algorithms (H)",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 11, 30, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
			},
			{
				Name:  "Lunch",
				Start: time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
				End:   time.Date(now.Year(), now.Month(), now.Day(), 13, 05, 0, 0, time.Local),
			},
		}, []Task{
			{
				Name:     "DSA Assignment",
				Duration: 120,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+2, 8, 0, 0, 0, time.Local),
				Priority: PRIORITY_IMPORTANT,
			},
			{
				Name:     "Side Project 1 Deadline: Rewrite Part 1 / 3",
				Duration: 60,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+14, 8, 0, 0, 0, time.Local),
				Priority: PRIORITY_UNIMPORTANT,
			},
			{
				Name:     "Side Project 1 Deadline: Rewrite Part 2 / 3",
				Duration: 60,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+14, 8, 0, 0, 0, time.Local),
				Priority: PRIORITY_UNIMPORTANT,
			},
			{
				Name:     "Side Project 1 Deadline: Rewrite Part 3 / 3",
				Duration: 60,
				Deadline: time.Date(now.Year(), now.Month(), now.Day()+14, 8, 0, 0, 0, time.Local),
				Priority: PRIORITY_UNIMPORTANT,
			},
		}
}

func typicalWeek() Input {
	now := time.Now()
	var events []Event
	var tasks []Task

	for i := range 7 {
		sleepStart := time.Date(now.Year(), now.Month(), now.Day()+i-1, 21, 0, 0, 0, time.Local)
		if i == 0 {
			sleepStart = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		}
		sleepEnd := time.Date(now.Year(), now.Month(), now.Day()+i, 6, 30, 0, 0, time.Local)
		events = append(events, Event{
			Name:  "Sleep",
			Start: sleepStart,
			End:   sleepEnd,
		})

		today := time.Date(now.Year(), now.Month(), now.Day()+i, 0, 0, 0, 0, time.Local)
		var schoolEvents []Event
		var schoolTasks []Task
		if i < 5 {
			if i%2 == 0 {
				schoolEvents, schoolTasks = aDay(today)
			} else {
				schoolEvents, schoolTasks = bDay(today)
			}
		}
		events = append(events, schoolEvents...)
		tasks = append(tasks, schoolTasks...)

		events = append(events, Event{
			Name:  "Dinner",
			Start: time.Date(now.Year(), now.Month(), now.Day()+i, 18, 0, 0, 0, time.Local),
			End:   time.Date(now.Year(), now.Month(), now.Day()+i, 18, 15, 0, 0, time.Local),
		})
	}

	return Input{
		Now:    now,
		Events: events,
		Tasks:  tasks,
	}
}

func TestSchedule(t *testing.T) {
	blocks, errs := Schedule(typicalWeek())
	if len(errs) > 0 {
		t.Fatal(errors.Join(errs...))
	}

	renderedBlocks := make([]templates.TimeBlock, len(blocks))
	for i, b := range blocks {
		var name string
		if b.Event != nil {
			name = b.Event.Name
		} else {
			name = b.Task.Name
		}
		renderedBlocks[i] = templates.TimeBlock{
			Name:  name,
			Start: b.Start,
			End:   b.End,
		}
	}

	marshalled, err := json.Marshal(renderedBlocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(marshalled))
}
