package scheduling

import (
	"algo-schedule/internal/templates"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

func aDay(now time.Time) []Reservable {
	return []Reservable{
		Event(
			"AP Statistics",
			time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, time.Local),
		),
		Event(
			"AP Microeconomics",
			time.Date(now.Year(), now.Month(), now.Day(), 9, 25, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, time.Local),
		),
		Event(
			"Multi-Variable Calculus (H)",
			time.Date(now.Year(), now.Month(), now.Day(), 10, 55, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, time.Local),
		),
		Event(
			"Lunch",
			time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 12, 30, 0, 0, time.Local),
		),
		Event(
			"World Religions",
			time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 14, 0, 0, 0, time.Local),
		),
		Task(
			"AP Stats Webassign",
			PRIORITY_IMPORTANT,
			60,
			time.Date(now.Year(), now.Month(), now.Day()+2, 8, 0, 0, 0, time.Local),
		),
		Task(
			"Multi Webassign",
			PRIORITY_IMPORTANT,
			90,
			time.Date(now.Year(), now.Month(), now.Day()+2, 8, 0, 0, 0, time.Local),
		),
		Task(
			"Econ Assignment: Part 1 / 2",
			PRIORITY_IMPORTANT,
			90,
			time.Date(now.Year(), now.Month(), now.Day()+7, 8, 0, 0, 0, time.Local),
		),
		Task(
			"Econ Assignment: Part 2 / 2",
			PRIORITY_IMPORTANT,
			90,
			time.Date(now.Year(), now.Month(), now.Day()+7, 8, 0, 0, 0, time.Local),
		),
	}
}

func bDay(now time.Time) []Reservable {
	return []Reservable{
		Event(
			"AP Physics",
			time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, time.Local),
		),
		Event(
			"Philosophy in Literature (H)",
			time.Date(now.Year(), now.Month(), now.Day(), 9, 25, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, time.Local),
		),
		Event(
			"Data Structures and Algorithms (H)",
			time.Date(now.Year(), now.Month(), now.Day(), 11, 30, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
		),
		Event(
			"Lunch",
			time.Date(now.Year(), now.Month(), now.Day(), 12, 45, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day(), 13, 05, 0, 0, time.Local),
		),
		Task(
			"DSA Assignment",
			PRIORITY_IMPORTANT,
			120,
			time.Date(now.Year(), now.Month(), now.Day()+2, 8, 0, 0, 0, time.Local),
		),
		Task(
			"Side Project 1 Deadline: Rewrite Part 1 / 3",
			PRIORITY_UNIMPORTANT,
			60,
			time.Date(now.Year(), now.Month(), now.Day()+14, 8, 0, 0, 0, time.Local),
		),
		Task(
			"Side Project 1 Deadline: Rewrite Part 2 / 3",
			PRIORITY_UNIMPORTANT,
			60,
			time.Date(now.Year(), now.Month(), now.Day()+14, 8, 0, 0, 0, time.Local),
		),
		Task(
			"Side Project 1 Deadline: Rewrite Part 3 / 3",
			PRIORITY_UNIMPORTANT,
			60,
			time.Date(now.Year(), now.Month(), now.Day()+14, 8, 0, 0, 0, time.Local),
		),
	}
}

func typicalWeek() Input {
	now := time.Now()
	var reservables []Reservable

	for i := range 7 {
		sleepStart := time.Date(now.Year(), now.Month(), now.Day()+i-1, 21, 0, 0, 0, time.Local)
		if i == 0 {
			sleepStart = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		}
		sleepEnd := time.Date(now.Year(), now.Month(), now.Day()+i, 6, 30, 0, 0, time.Local)
		reservables = append(reservables, Event(
			"Sleep",
			sleepStart,
			sleepEnd,
		))

		today := time.Date(now.Year(), now.Month(), now.Day()+i, 0, 0, 0, 0, time.Local)
		var schoolReserve []Reservable
		if i < 5 {
			if i%2 == 0 {
				schoolReserve = aDay(today)
			} else {
				schoolReserve = bDay(today)
			}
		}
		reservables = append(reservables, schoolReserve...)

		reservables = append(reservables, Event(
			"Dinner",
			time.Date(now.Year(), now.Month(), now.Day()+i, 18, 0, 0, 0, time.Local),
			time.Date(now.Year(), now.Month(), now.Day()+i, 18, 15, 0, 0, time.Local),
		))
	}

	return Input{
		Now:         now,
		Reservables: reservables,
	}
}

func TestSchedule(t *testing.T) {
	blocks, errs := Schedule(typicalWeek())
	if len(errs) > 0 {
		t.Fatal(errors.Join(errs...))
	}

	renderedBlocks := make([]templates.TimeBlock, len(blocks))
	for i, b := range blocks {
		renderedBlocks[i] = templates.TimeBlock{
			Name:  b.Reservable.Name,
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
