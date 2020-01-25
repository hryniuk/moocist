package coursera

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
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

func getCourseSyllabus(url string) CourseSyllabus {
	cs := CourseSyllabus{}

	c := colly.NewCollector(
		colly.AllowedDomains("coursera.org", "www.coursera.org"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"),
		colly.CacheDir("./coursera_cache"),
	)

	c.OnHTML("div.SyllabusModule", func(e *colly.HTMLElement) {
		week := Week{}

		e.ForEach("h1", func(_ int, f *colly.HTMLElement) {
			if week.Title != "" {
				log.Error("setting week title more than once")
			}
			week.Title = f.Text
		})

		e.ForEach("h2", func(_ int, f *colly.HTMLElement) {
			if week.Title != "" {
				log.Error("setting week title more than once")
			}
			week.Title = f.Text
		})

		e.ForEach("div.SyllabusModuleDetails div.items div", func(_ int, g *colly.HTMLElement) {
			if !strings.HasPrefix(g.Attr("class"), "Box") {
				dom := g.DOM
				duration := dom.Find("span span").Text()
				title := strings.TrimSuffix(dom.Text(), duration)
				task := Task{Title: title, Duration: duration}
				log.Debugf("adding task %s %s", title, duration)

				week.Tasks = append(week.Tasks, task)
			}
		})

		cs.Weeks = append(cs.Weeks, week)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	return cs
}
