package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hryniuk/moocist/pkg/coursera"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide course URL!")
	}

	courseURL := os.Args[1]

	fmt.Println(coursera.GetCourseInfo(courseURL))
}
