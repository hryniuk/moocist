package mooc

import (
	"testing"
	"time"
)

var tasksCount int
var cs CourseSyllabus

func init() {
	tasksCount = 6
	cs = CourseSyllabus{
		Weeks: []Week{
			Week{
				Title: "week 1",
				Items: []Item{
					Item{Title: "item 1"},
					Item{Title: "item 2"},
				},
			},
			Week{
				Title: "week 2",
				Items: []Item{
					Item{Title: "item 3"},
					Item{Title: "item 4"},
				},
			},
		},
	}
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func TestNextDayCalucatedCorrectly(t *testing.T) {
	withWeekendsTestCases := [][]time.Time{
		{date(2019, time.October, 10), date(2019, time.October, 11)},
		{date(2019, time.April, 20), date(2019, time.April, 21)},
		{date(2019, time.January, 6), date(2019, time.January, 7)},
		{date(2020, time.February, 28), date(2020, time.February, 29)},
	}

	withoutWeekendsTestCases := [][]time.Time{
		{date(2019, time.October, 10), date(2019, time.October, 11)},
		{date(2019, time.April, 20), date(2019, time.April, 22)},
		{date(2019, time.January, 6), date(2019, time.January, 7)},
		{date(2020, time.February, 28), date(2020, time.March, 2)},
	}

	skipWeekends := false
	for _, tc := range withWeekendsTestCases {
		nd := nextDay(tc[0], skipWeekends)
		if nd != tc[1] {
			t.Errorf("wrong next day expected %v got %v", tc[1], nd)
		}
	}

	skipWeekends = true
	for _, tc := range withoutWeekendsTestCases {
		nd := nextDay(tc[0], skipWeekends)
		if nd != tc[1] {
			t.Errorf("wrong next day expected %v got %v", tc[1], nd)
		}
	}
}

func TestPriorityIsSetCorrectly(t *testing.T) {
	expectedPriority := PriorityYellow
	opt := ExportOptions{
		TaskPriority: expectedPriority,
		AutoDate:     true,
	}
	exporter := TodoistExporter{Opt: opt}

	tasks := exporter.toTasks(cs)

	if len(tasks) != tasksCount {
		t.Errorf("incorrect tasks count, expected %d got %d", tasksCount, len(tasks))
	}

	for _, task := range tasks {
		if task.Type == Regular && task.Priority != expectedPriority {
			t.Errorf("incorrect regular task priority, expected %d got %d", expectedPriority, task.Priority)
		}

		if task.Type == TopLevel && task.Priority != PriorityNone {
			t.Errorf("incorrect top level task priority, expected %d got %d", PriorityNone, task.Priority)
		}
	}
}

func checkDates(t *testing.T, tasks []Task, expectedDates []time.Time) {
	taskDateIndex := 0
	for i, task := range tasks {
		if task.Type == Regular {
			if task.Date.Year() != expectedDates[taskDateIndex].Year() {
				t.Errorf("incorrect %d task date year, expected %d got %d", i, expectedDates[taskDateIndex].Year(), task.Date.Year())
			}
			if task.Date.Month() != expectedDates[taskDateIndex].Month() {
				t.Errorf("incorrect %d task date month, expected %d got %d", i, expectedDates[taskDateIndex].Month(), task.Date.Month())
			}
			if task.Date.Day() != expectedDates[taskDateIndex].Day() {
				t.Errorf("incorrect %d task date day, expected %d got %d", i, expectedDates[taskDateIndex].Day(), task.Date.Day())
			}
			taskDateIndex++
		}
	}
}

func TestDatesAreSetCorrectlyIncludingWeekends(t *testing.T) {
	opt := ExportOptions{
		StartingDate: time.Date(2019, time.October, 10, 0, 0, 0, 0, time.UTC),
		TasksPerDay:  uint32(1),
		SkipWeekends: false,
		AutoDate:     true,
	}
	exporter := TodoistExporter{Opt: opt}

	tasks := exporter.toTasks(cs)

	taskDates := []time.Time{
		date(2019, time.October, 10),
		date(2019, time.October, 11),
		date(2019, time.October, 12),
		date(2019, time.October, 13),
	}

	checkDates(t, tasks, taskDates)
}

func TestDatesAreSetCorrectlyWhenSkippingWeekends(t *testing.T) {
	opt := ExportOptions{
		StartingDate: time.Date(2019, time.October, 10, 0, 0, 0, 0, time.UTC),
		TasksPerDay:  uint32(1),
		SkipWeekends: true,
		AutoDate:     true,
	}
	exporter := TodoistExporter{Opt: opt}

	tasks := exporter.toTasks(cs)

	taskDates := []time.Time{
		date(2019, time.October, 10),
		date(2019, time.October, 11),
		date(2019, time.October, 14),
		date(2019, time.October, 15),
	}

	checkDates(t, tasks, taskDates)
}
