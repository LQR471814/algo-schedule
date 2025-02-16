package scheduling

import (
	"algo-schedule/internal/templates"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func aDay(now time.Time) []Event {
	return []Event{
		{
			Name:       "AP Statistics",
			Difficulty: DIFFICULTY_MEDIUM,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, time.Local),
		},
		{
			Name:       "AP Microeconomics",
			Difficulty: DIFFICULTY_MEDIUM,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 9, 25, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, time.Local),
		},
		{
			Name:       "Multi-Variable Calculus (H)",
			Difficulty: DIFFICULTY_LOW,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 10, 55, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, time.Local),
		},
		{
			Name:       "Lunch",
			Difficulty: DIFFICULTY_LOW,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 12, 30, 0, 0, time.Local),
		},
		{
			Name:       "World Religions",
			Difficulty: DIFFICULTY_LOW,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 14, 0, 0, 0, time.Local),
		},
	}
}

func bDay(now time.Time) []Event {
	return []Event{
		{
			Name:       "AP Physics",
			Difficulty: DIFFICULTY_MEDIUM,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, time.Local),
		},
		{
			Name:       "Philosophy in Literature (H)",
			Difficulty: DIFFICULTY_MEDIUM,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 9, 25, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, time.Local),
		},
		{
			Name:       "Data Structures and Algorithms (H)",
			Difficulty: DIFFICULTY_LOW,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 11, 30, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
		},
		{
			Name:       "Lunch",
			Difficulty: DIFFICULTY_LOW,
			Start:      time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day(), 13, 05, 0, 0, time.Local),
		},
	}
}

func typicalWeek(now time.Time) []Event {
	var events []Event

	for i := range 7 {
		events = append(events, Event{
			Name:       "Sleep",
			Difficulty: DIFFICULTY_LOW,
			Start:      time.Date(now.Year(), now.Month(), now.Day()+i-1, 21, 0, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day()+i, 6, 30, 0, 0, time.Local),
		})

		today := time.Date(now.Year(), now.Month(), now.Day()+i, 0, 0, 0, 0, time.Local)
		var schoolEvents []Event
		if i < 5 {
			if i%2 == 0 {
				schoolEvents = aDay(today)
			} else {
				schoolEvents = bDay(today)
			}
		}
		events = append(events, schoolEvents...)

		events = append(events, Event{
			Name:       "Dinner",
			Difficulty: DIFFICULTY_LOW,
			Start:      time.Date(now.Year(), now.Month(), now.Day()+i, 18, 0, 0, 0, time.Local),
			End:        time.Date(now.Year(), now.Month(), now.Day()+i, 18, 15, 0, 0, time.Local),
		})
	}
	events = append(events, Event{
		Name:       "Sleep",
		Difficulty: DIFFICULTY_LOW,
		Start:      time.Date(now.Year(), now.Month(), now.Day()+6, 21, 0, 0, 0, time.Local),
		End:        time.Date(now.Year(), now.Month(), now.Day()+7, 6, 30, 0, 0, time.Local),
	})

	return events
}

var testInputs []Input = []Input{
	{
		Events: typicalWeek(time.Now()),
		Quotas: []Quota{
			{
				Name:       "Read",
				Duration:   time.Minute * 30,
				Difficulty: DIFFICULTY_MEDIUM,
			},
		},
		Tasks: []Task{
			{
				Name:           "AP Stats Webassign",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 2,
			},
			{
				Name:           "Multi Webassign",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 2,
			},
			{
				Name:           "DSA Assignment",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_MEDIUM,
				TimeToDeadline: time.Hour * 24 * 2,
			},
			{
				Name:           "Econ Assignment: Part 1 / 2",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 8,
			},
			{
				Name:           "Econ Assignment: Part 2 / 2",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 8,
			},
			{
				Name:           "Bible Project: Part 1 / 6",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 7,
			},
			{
				Name:           "Bible Project: Part 2 / 6",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 7,
			},
			{
				Name:           "Bible Project: Part 3 / 6",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 7,
			},
			{
				Name:           "Bible Project: Part 4 / 6",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 7,
			},
			{
				Name:           "Bible Project: Part 5 / 6",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 7,
			},
			{
				Name:           "Bible Project: Part 6 / 6",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_HIGH,
				TimeToDeadline: time.Hour * 24 * 7,
			},
			{
				Name:           "Side Project 1 Deadline: Rewrite Part 1 / 3",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_MEDIUM,
				TimeToDeadline: time.Hour * 24 * 14,
			},
			{
				Name:           "Side Project 1 Deadline: Rewrite Part 1 / 3",
				Size:           SIZE_LARGE,
				Difficulty:     DIFFICULTY_MEDIUM,
				TimeToDeadline: time.Hour * 24 * 14,
			},
		},
	},
}

func TestSchedule(t *testing.T) {
	var blocks = make([]templates.TimeBlock, len(testInputs[0].Events))
	for i, e := range testInputs[0].Events {
		blocks[i] = templates.TimeBlock{
			Name:  e.Name,
			Start: e.Start,
			End:   e.End,
		}
	}

	marshalled, err := json.Marshal(blocks)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(marshalled))
}
