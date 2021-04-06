package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Disclaimer string                 `json:"disclaimer"`
	License    string                 `json:"license"`
	Timestamp  int                    `json:"timestamp"`
	Base       string                 `json:"base"`
	Rates      map[string]interface{} `json:"rates"`
}

func main() {
	resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=a8a051b461914baea0f9ae761060a4fe")
	if err != nil {
		fmt.Println("http get error:", err)
	}
	fmt.Println("resp:", resp.StatusCode)

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil ready error:", err)
	}
	var res Result
	json.Unmarshal(resData, &res)
	defer resp.Body.Close()
	fmt.Println(len(res.Rates))
	var sqlstr string
	for k, v := range res.Rates {
		sqlstr += fmt.Sprintf(`
		update cc set cc.ConvertToValue = %v from tbl_CurrencyConversion cc where cc.CurrencyCode = '%s' `, v, k)
	}
	fmt.Println("sqlstr:", sqlstr)
}
