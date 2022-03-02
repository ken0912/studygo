package main

import (
	"fmt"
	"regexp"
)

func main() {
	// // create a new collector
	// c := colly.NewCollector()

	// // authenticate
	// err := c.Post("http://example.com/login", map[string]string{"username": "admin", "password": "admin"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // attach callbacks after login
	// c.OnResponse(func(r *colly.Response) {
	// 	log.Println("response received", r.StatusCode)
	// })

	// // start scraping
	// c.Visit("https://example.com/")
	re := regexp.MustCompile(`\b(?i)(a|b)\b`)
	fmt.Println(re.FindAllString("nhooo.com a b", -1))
	fmt.Println(re.FindAllString("abc.org b", -1))
	fmt.Println(re.FindAllString("fb.com b a", -1))
}
