package coursera

func GetCourseInfo(url string) (CourseSyllabus, error) {
	return getCourseSyllabus(url), nil
}
