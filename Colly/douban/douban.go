package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type MovieInfo struct {
	Name     string `json:"name"`
	Director string `json:"director"`
	Type     string `json:"type"`
	Score    string `json:"Score"`
}

func main() {
	c := colly.NewCollector(
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./douban_cache"),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*douban.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second, // 两次请求 随机延迟5s 内
	})
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// start scaping the page under the link found
		e.Request.Visit(link)
	})

	Movies := make([]MovieInfo, 0, 200)
	fName := "Movies.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	c.OnHTML(`div[id=content]`, func(e *colly.HTMLElement) {
		movieinfo := MovieInfo{
			Name:     e.ChildText(`h1 > span[property="v:itemreviewed"]`),
			Director: e.ChildText(`span[class="attrs"] > a[rel="v:directedBy"]`),
			Type:     e.ChildText(`span[property="v:genre"]`),
			Score:    e.ChildText(`div[class="rating_self clearfix"] > strong`),
		}
		if movieinfo.Name != "" {
			Movies = append(Movies, movieinfo)
			fmt.Println(movieinfo)
		}

		// e.ForEach(`span[property="v:genre"]`, func(_ int, el *colly.HTMLElement) {
		// 	fmt.Println("typed:", el.Text)
		// })
	})

	// 访问初始 URL
	startUrl := "https://movie.douban.com/"
	_ = c.Visit(startUrl)

	// 等待爬虫结束
	c.Wait()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(Movies)
}
