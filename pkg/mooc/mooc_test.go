package mooc

import "testing"

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

func TestPriorityIsSetUpCorrectly(t *testing.T) {
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
