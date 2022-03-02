package main

import (
	"ESPopulateIdx/db"
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	elastic "github.com/olivere/elastic/v7"
)

var rowChan = make(chan map[string]interface{}, 1024)

type JobPosting struct {
	JobPostingID    string `json:"jobpostingid"`
	JobPostingTitle string `json:"jobpostingtitle"`
	StateCode       string `json:"statecode"`
}

func BulkeBuildIndex(stop chan bool) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://salarycnlinux01:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("client:", client)
	fmt.Println("connect to es success")
	fmt.Println("bulkBuildIndex Starts:", time.Now().Format("2006-01-02 15:04:05"))
	for row := range rowChan {
		_, err := client.Index().
			Index("jobboard").
			// Type("jobposting").
			Id(row["JobPostingID"].(string)).
			BodyJson(row).
			Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		// fmt.Printf("Indexed jobposting %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
	stop <- true
}

func GetDataFromMssql() {
	//connect mssql
	mssql := db.NewMssqlClient()
	db, err := mssql.Open()
	if err != nil {
		fmt.Println("mssql.Open() error:", err)
		return
	}
	defer db.Close()

	sqlstr := `
	select TOP 1000 jp.JobPostingID as JobPostingID, 
		isnull(jp.PostDate, '') as PostDate, 
		isnull(jp.JobPostingTitle, '') as JobPostingTitle, 
		isnull(jp.SEOFriendlyJobPostingTitle, '') as SEOTitle, 
		isnull(jp.CompanyID, '') as CompanyID, 
		isnull(jp.CompanyName, '') as CompanyName, 
		isnull(jp.SEOFriendlyCompanyName, '') as SEOCompanyName, 
		isnull(jp.JobLocationMetroCode, '') as MetroCode, 
		isnull(jp.JobLocationCity, '') as City, 
		isnull(jp.JobLocationStateCode, '') as StateCode, 
		isnull(jp.IsWorkFromHome, 0) as IsWorkFromHome 
	from tbl_jobposting(nolock) jp
	`
	rows, err := db.Query(sqlstr)

	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	columnLength := len(columns)

	cache := make([]interface{}, columnLength)
	for index, _ := range cache { //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}

	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		rowChan <- item
	}

	close(rowChan)
}
func query() {
	var res *elastic.SearchResult
	var err error
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://salarycnlinux01:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("client:", client)
	fmt.Println("connect to es success")
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("JobPostingTitle", "Manager"))
	boolQ.Must(elastic.NewMatchQuery("StateCode", "MA"))

	res, err = client.Search("jobboard").Query(boolQ).Size(3).From(0).Do(context.Background())
	printEmployee(res, err)

}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ JobPosting
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(JobPosting)
		fmt.Printf("%#v\n", t)
	}
	fmt.Println("res.TookInMillis:", res.TookInMillis, res.TotalHits())
}
func main() {

	// stop := make(chan bool)
	// go GetDataFromMssql()
	// go BulkeBuildIndex(stop)
	// for {
	// 	select {
	// 	case <-stop:
	// 		fmt.Println("bulkBuildIndex Done!", time.Now().Format("2006-01-02 15:04:05"))
	// 		return
	// 	default:
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }

	query()
	w := time.Now().Weekday()
	fmt.Println("w:", w)

}
