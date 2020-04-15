package coursera

import (
	"strings"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"github.com/hryniuk/moocist/pkg/mooc"
)

func getCourseSyllabus(url string) (mooc.CourseSyllabus, error) {
	cs := mooc.CourseSyllabus{}

	c := colly.NewCollector(
		colly.AllowedDomains("coursera.org", "www.coursera.org"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"),
	)

	c.OnHTML("div.SyllabusModule", func(e *colly.HTMLElement) {
		week := mooc.Week{}

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
				item := mooc.Item{Title: title, Duration: duration}
				log.Debugf("adding task %s %s", title, duration)

				week.Items = append(week.Items, item)
			}
		})

		cs.Weeks = append(cs.Weeks, week)
	})

	c.Visit(url)

	return cs, cs.Validate()
}
