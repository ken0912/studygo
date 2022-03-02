package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		// colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		colly.MaxDepth(2),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// ParseDouban(r)
		ExtractionData(r)
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnScraped(func(r *colly.Response) {

		fmt.Println("Finished", r.Request.URL)
	})

	// c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Visit("https://www.ncsl.org/research/labor-and-employment/state-minimum-wage-chart.aspx")
}

func ParseDouban(r *colly.Response) {
	// r.Save(r.FileName())
	body := string(r.Body)
	re_movieName := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
	re_rating := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	re_post := regexp.MustCompile(`<img width="100" alt="(.*?)" src="(.*?)"`)
	movieName := re_movieName.FindAllStringSubmatch(body, -1)
	rating := re_rating.FindAllStringSubmatch(body, -1)
	post := re_post.FindAllStringSubmatch(body, -1)
	for i, v := range movieName {
		fmt.Printf("MovieName:%s Rating:%s Post:%s \n", v[1], rating[i][1], post[i][2])
	}
}

func ExtractionData(r *colly.Response) {
	body := string(r.Body)
	body = trimHtml(body)
	// fmt.Println("body:", body)
	// body = `The new minimum wage varies across the state based on geographical location and, in New York City, employer size.`
	re := regexp.MustCompile(`(.*?)minimum wage(.*?)New York(.*?)[.?!]\s`)
	data := re.FindAllStringSubmatch(body, -1)
	for _, v := range data {
		fmt.Println("v:", v[0])
	}
}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
