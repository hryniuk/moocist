package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hryniuk/moocist/pkg/api"
	"github.com/hryniuk/moocist/pkg/coursera"
	"github.com/hryniuk/moocist/pkg/mooc"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\t%s <course URL or slug>\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	addr := flag.String("addr", "", "HTTP service address, when set, ignores the rest of options")
	priority := flag.Int("priority", int(mooc.PriorityNone), "priority of generated tasks 4 (grey) - 1 (red)")

	flag.Parse()

	if len(*addr) > 0 {
		s := api.NewServer()
		s.Run(*addr)
		return
	}

	courseRef := flag.Args()[0]
	courseSlug, err := mooc.GetMOOCSlug(courseRef)
	if err != nil {
		log.Printf("cannot retrieve course slug: %s\n", err)
		os.Exit(1)
	}

	importer := coursera.SlugImporter{Slug: courseSlug}
	syllabus, err := importer.Import()
	if err != nil {
		log.Printf("cannot download course syllabus: %s\n", err)
		os.Exit(1)
	}

	taskPriority := mooc.PriorityNone
	if 1 <= *priority && *priority <= 4 {
		taskPriority = mooc.Priority(*priority)
	}

	opt := mooc.ExportOptions{TaskPriority: mooc.Priority(taskPriority)}
	exporter := mooc.TodoistExporter{Opt: opt}
	b, err := exporter.Export(syllabus)
	if err != nil {
		log.Printf("cannot export to CSV file: %s\n", err)
		os.Exit(1)
	}

	filename := fmt.Sprintf("%s.csv", courseSlug)
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("cannot create %s file: %s\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(string(b))
	if err != nil {
		log.Printf("cannot write course template to %s file: %s\n", filename, err)
		os.Exit(1)
	}
}
