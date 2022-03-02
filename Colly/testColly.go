package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// fmt.Println(os.Args[1:])
	// fmt.Println("time:", time.Now().UTC())
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
	// testcollydom()
	/*
		resp, err := http.Get(`https://www.xe.com/currencyconverter/convert/?Amount=1&From=ZWD&To=USD`)
		if err != nil {
			log.Println("http Get err:", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ioutil ReadAll err:", err)
		}
		// fmt.Println("body:", string(body))
		if err = ioutil.WriteFile("xe_http.html", body, 0666); err != nil {
			fmt.Println("WriteFile Error:", err)
			return
		}
	*/
}

type address struct {
	AddressCountry  string `json:"addressCountry"`
	AddressLocality string `json:"addressLocality"`
	AddressRegion   string `json:"addressRegion"`
	PostalCode      string `json:"postalCode"`
	StreetAddress   string `json:"streetAddress"`
}

func testcollydom() {
	c := colly.NewCollector(
		colly.Async(true),
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
		r.Save("xecom.html")
		fmt.Println("Visited", r.Request.URL)
	})

	// c.OnHTML("body", GetSalaryDesc)
	c.OnHTML("div[class='unit-rates___StyledDiv-sc-1dk593y-0 dEqdnx']", GetXe)
	// c.OnHTML(".paginator a", func(_ int, e *colly.HTMLElement) {
	// 	fmt.Println("url:", e.Attr("href"))
	// 	e.Request.Visit(e.Attr("href"))
	// })

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// c.Visit("https://www.salary.com/" + "about-us")
	c.Visit("https://www.xe.com/currencyconverter/convert/?Amount=1&From=ZWD&To=USD")
	// c.Visit("https://www.indeed.com/cmp/Salary.com")
	// c.Visit("https://www.indeed.com/cmp/Google")
	c.Wait()

}
func GetXe(e *colly.HTMLElement) {
	value := strings.Split(e.Text, "=")
	fmt.Println("xe:", value)
	for _, v := range value {
		fmt.Println(strings.Split(strings.TrimSpace(v), " ")[0])
	}

}

func GetSalaryDesc(e *colly.HTMLElement) {
	// log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
	// 	strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
	var Desc string
	e.ForEachWithBreak("p", func(_ int, elem *colly.HTMLElement) bool {
		Desc = Desc + elem.Text
		if len(Desc) > 1500 {
			return false
		}
		return true
	})
	// fmt.Println("Desc length:", len(Desc))
	// fmt.Println("Desc:", Desc)
	var mapResult map[string]interface{}
	e.ForEachWithBreak("script[type='application/ld+json']", func(_ int, elem *colly.HTMLElement) bool {
		// fmt.Println("json data:", elem.Text)
		if err := json.Unmarshal([]byte(elem.Text), &mapResult); err != nil {
			fmt.Println("err:", err)
			return false
		}
		return true
	})
	fmt.Println("mapResult:", mapResult["address"])

}
