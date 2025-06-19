package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
)

func scrape(url string) (to_hash string) {
	//Strips protocol so that the domain is solely the IP/Port combo. Otherwise colly will not scrape web page
	domain := strings.Trim(url,"http://")
	

	c:= colly.NewCollector(
		colly.AllowedDomains(
			domain,
		),
	)

	c.OnHTML("h3", func(e *colly.HTMLElement){
		to_hash = e.Text
		fmt.Println("To Hash:",to_hash)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response){
		cookies := c.Cookies(r.Request.URL.String())
		fmt.Println(cookies)
	})

	c.Visit(url)
	
	return
}


func main() {

	url := os.Args[1]
	fmt.Println(scrape(url))

}