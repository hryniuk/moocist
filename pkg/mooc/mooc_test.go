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

func TestPriorityIsSetCorrectly(t *testing.T) {
	expectedPriority := PriorityYellow
	opt := ExportOptions{
		TaskPriority: expectedPriority,
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
	}
	exporter := TodoistExporter{Opt: opt}

	tasks := exporter.toTasks(cs)

	taskDates := []time.Time{
		time.Date(2019, time.October, 10, 0, 0, 0, 0, time.UTC),
		time.Date(2019, time.October, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2019, time.October, 12, 0, 0, 0, 0, time.UTC),
		time.Date(2019, time.October, 13, 0, 0, 0, 0, time.UTC),
	}

	checkDates(t, tasks, taskDates)
}

func TestDatesAreSetCorrectlyWhenSkippingWeekends(t *testing.T) {
	opt := ExportOptions{
		StartingDate: time.Date(2019, time.October, 10, 0, 0, 0, 0, time.UTC),
		TasksPerDay:  uint32(1),
		SkipWeekends: true,
	}
	exporter := TodoistExporter{Opt: opt}

	tasks := exporter.toTasks(cs)

	taskDates := []time.Time{
		time.Date(2019, time.October, 10, 0, 0, 0, 0, time.UTC),
		time.Date(2019, time.October, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2019, time.October, 14, 0, 0, 0, 0, time.UTC),
		time.Date(2019, time.October, 15, 0, 0, 0, 0, time.UTC),
	}

	checkDates(t, tasks, taskDates)
}
