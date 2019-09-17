package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"sync"

	"github.com/gocolly/colly"
)

type product struct {
	ProductURL string
	Name       string
	Price      string
	ID         string
}

var searchURL = []string{"https://www.amazon.co.uk/s?k=", "https://www.amazon.fr/s?k=",
	"https://www.amazon.es/s?k=", "https://www.amazon.it/s?k=",
	"https://www.amazon.de/s?k="}

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.co.uk", "www.amazon.de", "www.amazon.es", "www.amazon.fr", "www.amazon.it"),
	)
	products := make(map[string][]product)
	searchTerm := "ps4 pro"
	counter := 0

	var a []product

	// here for each element inside the selector in the on html function we are lookin for a product
	c.OnHTML("span[data-component-type=\"s-search-results\"]", func(e *colly.HTMLElement) {
		e.ForEach("div.s-result-list.sg-row > div", func(index int, el *colly.HTMLElement) {
			name := el.DOM.Find("span.a-size-medium.a-color-base.a-text-normal").Text()
			price := el.DOM.Find("div.a-section.a-spacing-none.a-spacing-top-small span.a-price > span.a-offscreen").Text()
			urls := el.ChildAttr("a.a-link-normal.a-text-normal", "href")
			id := el.Attr("data-asin")
			product := product{Name: name, Price: price, ProductURL: urls, ID: id}
			if v, found := products[id]; found {
				products[id] = append(v, product)
			} else {
				products[id] = append(a, product)

			}

		})

	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnScraped(func(r *colly.Response) {
		// on scraped will be called once the onHtml fuction is completed
		fmt.Println("visited")
		// a counter that increments so we know that we already finished all the iterations in the loop
		counter++
		if counter == len(searchURL) {
			fmt.Println("creating json")
			// converting the product map into json
			out, err := json.Marshal(products)
			if err != nil {
				panic(err)
			}
			// output the file into json file
			err = ioutil.WriteFile("output.json", out, 0644)
		}

	})

	var wg sync.WaitGroup
	// Tell the 'wg' WaitGroup how many threads/goroutines
	//   that are about to run concurrently.
	wg.Add(len(searchURL))
	for i := 0; i < len(searchURL); i++ {
		// Spawn a thread for each iteration in the loop.
		// Pass 'i' into the goroutine's function
		//   in order to make sure each goroutine
		//   uses a different value for 'i'.
		go func(t int) {

			fmt.Println(searchURL[t] + url.QueryEscape(searchTerm))
			c.Visit(searchURL[t] + url.QueryEscape(searchTerm))
			// At the end of the goroutine, tell the WaitGroup
			//   that another thread has completed.
			defer wg.Done()

		}(i)

	}
	// Wait for `wg.Done()` to be exectued the number of times
	//   specified in the `wg.Add()` call.
	// `wg.Done()` should be called the exact number of times
	//   that was specified in `wg.Add()`.
	// If the numbers do not match, `wg.Wait()` will either
	//   hang infinitely or throw a panic error.
	wg.Wait()
	fmt.Println("Finished for loop")
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
