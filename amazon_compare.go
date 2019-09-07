package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	ProductUrl string
	Name       string
	Price      string
}

func main() {

	// var name string
	// flag.StringVar(&name, "opt", "", "Usage")

	flag.Parse()
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.co.uk", "www.amazon.de", "www.amazon.es", "www.amazon.fr", "www.amazon.it"),
	)
	m := make(map[string][]item)
	var a []item
	c.OnHTML("span[data-component-type=\"s-search-results\"]", func(e *colly.HTMLElement) {
		e.ForEach("div.s-result-list.sg-row > div", func(index int, el *colly.HTMLElement) {
			name := el.DOM.Find("span.a-size-medium.a-color-base.a-text-normal").Text()
			price := el.DOM.Find("div.a-section.a-spacing-none.a-spacing-top-small span.a-price > span.a-offscreen").Text()
			urls := el.ChildAttr("a.a-link-normal.a-text-normal", "href")
			id := el.Attr("data-asin")
			product := item{Name: name, Price: price, ProductUrl: urls}
			if v, found := m[id]; found {
				m[id] = append(v, product)
			} else {
				m[id] = append(a, product)

			}

		})

	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("result.csv")
		checkError("Cannot create file", err)
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		for _, value := range m {
			out, err := json.Marshal(value)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(out))
			var strArray []string
			strArray = append(strArray, string(out))
			err = writer.Write(strArray)
			checkError("Cannot write to file", err)
		}
		fmt.Println(m)
	})

	go c.Visit("https://www.amazon.co.uk/s?k=ps4+pro&ref=nb_sb_noss_1")
	go c.Visit("https://www.amazon.de/s?k=ps4+pro&ref=nb_sb_noss_1")
	go c.Visit("https://www.amazon.es/s?k=ps4+pro&__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&ref=nb_sb_noss_1")
	go c.Visit("https://www.amazon.fr/s?k=ps4+pro&__mk_fr_FR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&ref=nb_sb_noss_1")
	c.Visit("https://www.amazon.it/s?k=ps4+pro&__mk_it_IT=%C3%85M%C3%85%C5%BD%C3%95%C3%91&ref=nb_sb_noss_1")

}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
