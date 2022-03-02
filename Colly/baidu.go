package main

import (
	"fmt"

	"github.com/crawlab-team/crawlab-go-sdk/entity"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	// 生成 colly 采集器
	c := colly.NewCollector(
		colly.AllowedDomains("www.baidu.com"),
		// colly.Async(true),
		// colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "http://www.sse.com.cn")
		//关键头 如果没有 则返回 错误
		r.Headers.Set("Referer", "http://www.sse.com.cn/assortment/stock/list/share/")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	})
	// 抓取结果数据钩子函数
	c.OnHTML(".result.c-container", func(e *colly.HTMLElement) {
		// 抓取结果实例
		item := entity.Item{
			"title": e.ChildText("h3.t > a"),
			"url":   e.ChildAttr("h3.t > a", "href"),
		}

		// 打印抓取结果
		fmt.Println(item)

		// 取消注释调用 Crawlab Go SDK 存入数据库
		//_ = crawlab.SaveItem(item)
	})

	// 分页钩子函数
	c.OnHTML("a.n", func(e *colly.HTMLElement) {

		_ = c.Visit("https://www.baidu.com" + e.Attr("href"))
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("r.Body", string(r.Body))
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	// 访问初始 URL
	startUrl := "https://www.baidu.com/s?wd=salary"
	_ = c.Visit(startUrl)

	// 等待爬虫结束
	c.Wait()
}
