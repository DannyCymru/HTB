package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
	"crypto/md5"
    "encoding/hex"
)

func scrape(url string, c *colly.Collector) (to_hash string, cookie string) {

	//Scrape webpage for the text we need to hash
	c.OnHTML("h3", func(r *colly.HTMLElement){
		to_hash = r.Text
	})

	//Collect the cookie that the application sets to ensure our hash gets accepted
	c.OnResponse(func(r *colly.Response){
		cookies := c.Cookies(r.Request.URL.String())
		cookie = cookies[0].String()
	})

	c.Visit(url)
	c.Wait()

	return
}

func hash(to_hash string)(md5_string string) {
	to_md5 := md5.Sum([]byte(to_hash))
	md5_string = hex.EncodeToString(to_md5[:])
	return 
}

func post(md5_string string, cookie string, url string, c *colly.Collector)(flag string){

	c.OnRequest(func(r *colly.Request){
		var id string 
		id = strings.Trim(cookie,"PHPSESSID=")
		r.Headers.Set("PHPSESSID", id)
	})


	c.OnHTML("p", func(r *colly.HTMLElement){
		flag = r.Text
	})

	c.Post(url, map[string]string{
		"hash": md5_string,
	})

	c.Wait()

	return
}

func main() {
	url := os.Args[1]

	//If block to ensure that whether you provide the IP with the protocol or not, that required variables get appropriately set.
	//Domain needs to solely be the IP/PortNumber while URL, needs the protocol.
	var domain string
	if  strings.Index(url,"http://") >=0 {
		domain = strings.Trim(url,"http://")
	} else {
		domain = url
		url = "http://" + domain
	}
	
	//Create collection config for our actions
	c:= colly.NewCollector(
		colly.AllowedDomains(
			domain,
		),
		colly.Async(true),
	)

	c.SetProxy("http://127.0.0.1:8080")

	var to_hash, cookie = scrape(url, c)

	fmt.Println("Text to hash:", to_hash)
	
	md5_string := hash(to_hash)

	fmt.Println("MD5 Hash:", md5_string)

	flag := post(md5_string, cookie, url, c)

	fmt.Println("Flag:", flag)
}