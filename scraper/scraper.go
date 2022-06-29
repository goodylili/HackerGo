package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"net/http"
	"time"
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

/*
RandomNewsPage scrapes a random news page using the ScrapeNews function
*/
func RandomNewsPage(page int) []Information {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(30-1) + 1
	return ScrapeNews(random)
}

// NumberOfJobs selects news to be displayed using the random or a specified page, It just controls the number of news to be displayed/*
func NumberOfJobs(number int) {
	jobs := ScrapeJobs()
	for index, structs := range jobs {
		if index <= number {
			format := fmt.Sprintf("%s \n %s \n\n", structs.Title, structs.Link)
			fmt.Println(format)
		}

	}

}

// NumberOfNews selects news to be displayed using the random or a specified page, It just controls the number of news to be displayed/*
func NumberOfNews(pages, number int) {
	jobs := ScrapeNews(pages)
	for index, structs := range jobs {
		if index <= number {
			format := fmt.Sprintf("%s \n %s \n\n", structs.Title, structs.Link)
			fmt.Println(format)
		}

	}

}

func selectRandom() {
	// select
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

	// 31642927

	return Info
}
