package mooc

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

type Importer interface {
	Import() (CourseSyllabus, error)
}

type Exporter interface {
	Export(CourseSyllabus) ([]byte, error)
}
