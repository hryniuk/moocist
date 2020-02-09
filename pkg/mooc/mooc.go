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
