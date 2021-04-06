package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type CurrencyConver struct {
	Scur   string  `json:"scur"`
	Tcur   string  `json:"tcur"`
	Rate   float64 `json:"rate"`
	Update string  `json:"update"`
}

func main() {
	scur := "USD"
	tcur := "ZWD"
	url := fmt.Sprintf("https://huobiduihuan.51240.com/?f=%s&t=%s&j=1", scur, tcur)
	rate, date := parseUrls(url)

	c := CurrencyConver{
		Scur:   scur,
		Tcur:   tcur,
		Rate:   rate,
		Update: date,
	}
	res, err := json.Marshal(&c)
	if err != nil {
		fmt.Println("json Marshal err:", err)
	}
	fmt.Println("res:", string(res))

}

func fetch(url string) string {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}
func parseUrls(url string) (float64, string) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	re := regexp.MustCompile(`<title>(.*?)</title>`)
	content := re.FindString(body)
	rerate := regexp.MustCompile(`\d+`)
	ratecontent := rerate.FindAllStringSubmatch(content, -1)
	// fmt.Println("ratecontent:", ratecontent)
	var rate float64
	lenr := len(ratecontent)
	if lenr > 1 {
		var r []string
		for _, v := range ratecontent[1:] {
			r = append(r, v[0])
		}
		ratef, err := strconv.ParseFloat(strings.Join(r, `.`), 64)
		if err != nil {
			fmt.Println("err:", err)
		}
		rate = ratef
	} else if lenr == 1 {
		rate, _ = strconv.ParseFloat(ratecontent[1:][0][0], 64)
	} else {
		rate = 0
	}

	//parse rate update date
	pattern, _ := regexp.Compile(`汇率更新时间：(.*?)）</div>`)
	date := pattern.FindStringSubmatch(body)[1]

	if rate == 0 {
		fmt.Println("err: invalid currency code!")
	}
	// fmt.Println("rate:", rate)
	return rate, date

}
