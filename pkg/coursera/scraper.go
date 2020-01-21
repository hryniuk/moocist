package coursera

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Task struct {
	Title    string
	Duration string
}

type Week struct {
	Title string
	Tasks []Task
}

type Course struct {
	Weeks []Week
}

func getCourseSyllabus(url string) string {
	c := colly.NewCollector(
		colly.AllowedDomains("coursera.org", "www.coursera.org"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"),
		colly.CacheDir("./coursera_cache"),
	)

	c.OnHTML("div.SyllabusModule", func(e *colly.HTMLElement) {
		e.ForEach("h1", func(_ int, f *colly.HTMLElement) {
			fmt.Println(f.Text)
		})

		e.ForEach("div.SyllabusModuleDetails div.items div", func(_ int, g *colly.HTMLElement) {
			if !strings.HasPrefix(g.Attr("class"), "Box") {
				dom := g.DOM
				time := dom.Find("span span").Text()
				title := strings.TrimSuffix(dom.Text(), time)
				fmt.Println(title)
				fmt.Println(time)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	return ""
}
