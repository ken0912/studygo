package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strings"
)

type pageHelp struct {
	Total int `json:"total"`
}

type pageResult struct {
	CompanyAbbr string `json:"COMPANY_ABBR"`
	CompanyCode string `json:"COMPANY_CODE"`
	EnglishAbbr string `json:"ENGLISH_ABBR"`
	ListingTime string `json:"LISTING_DATE"`
	MoveTime    string `json:"QIANYI_DATE"`
}

type PageResult struct {
	AreaName string `json:"areaName"`
	Page pageHelp `json:"pageHelp"`
	Result []pageResult `json:"result"`
	StockType string `json:"stockType"`
}

func (receiver Collector) ScrapeJs(url string) (error,[]*PageResult) {

	receiver.MLog.GetLogHandle().WithFields(log.Fields{"URL":url}).Info("URL...")

	c := colly.NewCollector(colly.UserAgent(RandomString()),colly.AllowURLRevisit())

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"

	res := make([]*PageResult, 0)

	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", RandomString())
		r.Headers.Set("Host", "query.sse.com.cn")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "http://www.sse.com.cn")
		//关键头 如果没有 则返回 错误
		r.Headers.Set("Referer", "http://www.sse.com.cn/assortment/stock/list/share/")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		receiver.MLog.GetLogHandle().WithFields(log.Fields{"Request":fmt.Sprintf("%+v",*r),"Headers":fmt.Sprintf("%+v",*r.Headers)}).Info("Begin Visiting...")
	})

	c.OnError(func(_ *colly.Response, err error) {
		receiver.MLog.GetLogHandle().WithFields(log.Fields{"error":err}).Info("Something went wrong:")
	})
	c.OnResponse(func(r *colly.Response) {
		receiver.MLog.GetLogHandle().WithFields(log.Fields{"Headers":r.Headers}).Info("Receive Header")
	})
	//scraped item from body
	//finish
	c.OnScraped(func(r *colly.Response) {
		var item  PageResult
		err := json.Unmarshal(r.Body, &item)
		if err != nil {
			receiver.MLog.GetLogHandle().WithFields(log.Fields{"err":err,"res":string(r.Body)}).Error("Receive Error ")
			if strings.Contains(string(r.Body),"error"){
				c.Visit(url)
			}
		}
		if len(item.Result) > 1{
			res = append(res,&item)
			receiver.MLog.GetLogHandle().WithFields(log.Fields{"item":item}).Info("Receive message ")
			return //结束递归
		}
		SubUrl :=  "http://query.sse.com.cn/security/stock/getStockListData2.do?stockType=1&pageHelp.cacheSize=1&pageHelp.beginPage=1&pageHelp.pageSize="
		//SubUrl += strconv.Itoa(item.Page.Total)
		SubUrl += '25'
		c.Visit(SubUrl)
	})

	c.Visit(url)
	return nil,res

}
func (receiver Collector) ScrapeJsTest()  error {
    //第一次只获取一条数据目的是为了获取数据总条数，下次递归拼接URL一次获取所有数据
	UrlA := "http://query.sse.com.cn/security/stock/getStockListData2.do?stockType=1&pageHelp.cacheSize=1&pageHelp.pageSize=1&pageHelp.beginPage=1"
	receiver.ScrapeJs(UrlA)
	return nil
}

func main() {
 var c Collector
 c.MLog=receiver.MLog
 c.ScrapeJsTest()
 return
}
