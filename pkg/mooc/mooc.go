package mooc

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Task struct {
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

type Week struct {
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

type CourseSyllabus struct {
	Weeks []Week `json:"weeks"`
}

func (cs CourseSyllabus) Validate() error {
	if len(cs.Weeks) == 0 {
		return errors.New("empty weeks")
	}

	for i, week := range cs.Weeks {
		if len(week.Title) == 0 {
			return fmt.Errorf("empty title in week with index %d", i)
		}

		if len(week.Tasks) == 0 {
			return fmt.Errorf("empty tasks in week with index %d", i)
		}

		for j, task := range week.Tasks {
			if len(task.Title) == 0 {
				return fmt.Errorf("empty title in task with index %d in week index %d", j, i)
			}

			if len(task.Duration) == 0 {
				return fmt.Errorf("empty duration in task with index %d in week index %d", j, i)
			}
		}
	}

	return nil
}

type Importer interface {
	Import() (CourseSyllabus, error)
}

type Exporter interface {
	Export(cs CourseSyllabus) ([]byte, error)
}

type JsonExporter struct {
}

func (e *JsonExporter) Export(cs CourseSyllabus) ([]byte, error) {
	return json.Marshal(cs)
}
