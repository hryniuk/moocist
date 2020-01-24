package main

import (
	"fmt"
	"os"

	"github.com/hryniuk/moocist/pkg/coursera"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	if len(os.Args) < 2 {
		log.Fatal("provide course URL!")
	}

	courseURL := os.Args[1]

	fmt.Println(coursera.GetCourseInfo(courseURL))
}
