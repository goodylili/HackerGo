package scraper

import (
	"fmt"
	"math/rand"
	"time"
)

/*
RandomNewsPage scrapes a random news page using the ScrapeNews function
*/

/*
AllFromPage prints all the news in a page8
*/
func AllFromPage(page int) {
	scrapePage := ScrapeNews(page)
	for _, structs := range scrapePage {
		format := fmt.Sprintf("%s \n %s\n", structs.Title, structs.Link)
		fmt.Println(format)
	}
}

/*
RandomNewsPage generate a random number for the random command
*/
func RandomNewsPage() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(30-1) + 1
	return random
}

// NumberOfJobs selects news to be displayed using the random or a specified page, It just controls the number of news to be displayed/*
func NumberOfJobs(number int) {
	jobs := ScrapeJobs()
	for index, structs := range jobs {
		if index <= number {
			format := fmt.Sprintf("%s \n %s\n", structs.Title, structs.Link)
			fmt.Println(format)
		}

	}

}

// NumberOfNews selects news to be displayed using the random or a specified page, It just controls the number of news to be displayed/*
func NumberOfNews(pages, number int) {
	jobs := ScrapeNews(pages)
	for index, structs := range jobs {
		if index <= number {
			format := fmt.Sprintf("%s \n %s\n", structs.Title, structs.Link)
			fmt.Println(format)
		}

	}

}

/*
AllJobsFromPage returns all the jobs in the page
*/
func AllJobsFromPage() {
	jobs := ScrapeJobs()
	for _, structs := range jobs {
		format := fmt.Sprintf("%s \n %s \n", structs.Title, structs.Link)
		fmt.Println(format)
	}
}
