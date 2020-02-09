package coursera

import (
	"fmt"

	"github.com/hryniuk/moocist/pkg/mooc"
)

const baseCourseURLFormat = "https://www.coursera.org/learn/%s#syllabus"

type SlugImporter struct {
	Slug string
}

func (i *SlugImporter) Import() (mooc.CourseSyllabus, error) {
	url := fmt.Sprintf(baseCourseURLFormat, i.Slug)
	return getCourseSyllabus(url)
}
