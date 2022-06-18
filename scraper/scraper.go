package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func scrapeItOff(page int) []Information {
	info := make([]Information, 0)
	webUrl, num := "https://news.ycombinator.com/news?p=", page
	fullURL := fmt.Sprintf("%s%d", webUrl, num)
	response, _ := http.Get(fullURL)

	document, _ := goquery.NewDocumentFromReader(response.Body)

	document.Find("tr.athing").Each(func(index int, selector *goquery.Selection) {
		//index = 5,
		title := selector.Find("td.title").Text()

		getLink, _ := selector.Find("a.titlelink").Attr("href")

		info = append(info, Information{
			Title: title,
			Link:  getLink,
		})

	})
	return info
}

func randomNewsPage(page int) {

	return
}

func selectRandom() {
	// select
}

func selectNewsNumber() {

}