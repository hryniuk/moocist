package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)

	}
	courseRef := flag.Args()[0]
	courseSlug, err := mooc.GetMOOCSlug(courseRef)
	if err != nil {
		log.Fatal(err)
	}

	importer := coursera.SlugImporter{Slug: courseSlug}
	syllabus, err := importer.Import()
	if err != nil {
		log.Fatal(err)
	}

	exporter := mooc.TodoistExporter{}
	b, err := exporter.Export(syllabus)
	if err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("%s.csv", courseSlug)
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(string(b))
	if err != nil {
		log.Fatal(err)
	}
}
