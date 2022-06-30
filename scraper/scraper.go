package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

/*
ScrapeNews Scrapes the website for other handlers to use based on the specified page
*/
func ScrapeNews(page int) []Information {
	var Info = make([]Information, 0)
	webUrl, num := "https://news.ycombinator.com/news?p=", page
	fullURL := fmt.Sprintf("%s%d", webUrl, num)
	response, _ := http.Get(fullURL)

	document, _ := goquery.NewDocumentFromReader(response.Body)

	document.Find("tr.athing").Each(func(index int, selector *goquery.Selection) {
		//index = 5,
		title := selector.Find("td.title").Text()

		getLink, _ := selector.Find("a.titlelink").Attr("href")

		Info = append(Info, Information{
			Title: title,
			Link:  getLink,
		})

	})
	return Info
}

func ScrapeJobs() []Information {
	var Info = make([]Information, 0)
	response, _ := http.Get("https://news.ycombinator.com/jobs")

	document, _ := goquery.NewDocumentFromReader(response.Body)

	document.Find("tr.athing").Each(func(index int, selector *goquery.Selection) {
		//index = 5,
		title := selector.Find("td.title").Text()

		getLink, _ := selector.Find("a.titlelink").Attr("href")

		Info = append(Info, Information{
			Title: title,
			Link:  getLink,
		})

	})

	return Info
}
