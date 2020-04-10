package mooc

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetSlugFromURLPath(path string) (string, error) {
	pathParts := strings.Split(path, "/")
	if len(pathParts) == 0 {
		return "", errors.New("cannot get slug from empty path")
	}

	// Assuming /learn/english-principles/...#... structure here.
	// Could be change to regex.
	for i, p := range pathParts {
		if p == "learn" && len(pathParts)-i >= 1 {
			return pathParts[i+1], nil
		}
	}

	return path, nil
}

func GetMOOCSlug(courseReference string) (string, error) {
	uriReference, err := url.Parse(courseReference)
	if err != nil {
		return "", err
	}

	return GetSlugFromURLPath(uriReference.Path)
}

type Item struct {
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

type Week struct {
	Title string `json:"title"`
	Items []Item `json:"items"`
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

		if len(week.Items) == 0 {
			return fmt.Errorf("empty tasks in week with index %d", i)
		}

		for j, task := range week.Items {
			if len(task.Title) == 0 {
				return fmt.Errorf("empty title in task with index %d in week index %d", j, i)
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

type Priority uint8

const (
	PriorityNone Priority = iota
	PriorityRed
	PriorityYellow
	PriorityBlue
	PriorityGrey
)

const (
	TopLevel = "week"
	Regular  = "task"
)

// Task is a Todoist equivalent of a MOOC "item":
//  * week info
//  * lecture video
//  * quiz etc.
type Task struct {
	Priority Priority
	No       uint32
	Title    string
	Type     string
	Date     time.Time
}

type ExportOptions struct {
	TaskPriority Priority
	StartingDate time.Time
	TasksPerDay  uint32
	SkipWeekends bool
	AutoDate     bool
}

type TodoistExporter struct {
	Opt ExportOptions
}

func isWeekendDay(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func nextDay(t time.Time, skipWeekends bool) time.Time {
	next := t

	next = next.AddDate(0, 0, 1)
	for skipWeekends && isWeekendDay(next) {
		next = next.AddDate(0, 0, 1)
	}

	return next
}

func (e *TodoistExporter) toTasks(cs CourseSyllabus) []Task {
	taskNo := uint32(0)
	var tasks []Task
	taskDate := e.Opt.StartingDate

	for _, week := range cs.Weeks {
		weekTask := Task{
			Priority: PriorityNone,
			No:       taskNo,
			Type:     TopLevel,
			Title:    week.Title,
		}
		tasks = append(tasks, weekTask)

		for _, item := range week.Items {
			itemTask := Task{
				Priority: e.Opt.TaskPriority,
				No:       taskNo,
				Title:    item.Title,
				Type:     Regular,
				Date:     taskDate,
			}
			if e.Opt.AutoDate {
				taskDate = nextDay(taskDate, e.Opt.SkipWeekends)
			}
			tasks = append(tasks, itemTask)
		}
	}

	return tasks
}

func taskToCSV(t Task) []string {
	dateFormat := "02/01/2006"
	priorityStr := strconv.Itoa(int(t.Priority))
	var defaultDate time.Time
	dateStr := ""

	if t.Date != defaultDate {
		dateStr = t.Date.Format(dateFormat)
	}
	taskType := "task"
	if t.Type == TopLevel {
		taskType = "section"
	}

	return []string{taskType, t.Title, priorityStr, "", "", "", dateStr, "", ""}
}

func (e *TodoistExporter) Export(cs CourseSyllabus) ([]byte, error) {
	header := []string{"TYPE", "CONTENT", "PRIORITY", "INDENT", "AUTHOR", "RESPONSIBLE", "DATE", "DATE_LANG", "TIMEZONE"}
	records := [][]string{
		header,
	}

	tasks := e.toTasks(cs)
	for _, task := range tasks {
		records = append(records, taskToCSV(task))
	}

	b := bytes.NewBuffer([]byte{})
	w := csv.NewWriter(b)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
		return []byte{}, err
	}
	return b.Bytes(), nil
}
