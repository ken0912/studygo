package main

import (
	"crypto/tls"
	"fmt"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/gomail.v2"
)

type AutoCodeStruct struct {
	StructName string `json:"structName"`
	TableName  string `json:"tableName"`
	// PackageName        string  `json:"packageName"`
	// Abbreviation       string  `json:"abbreviation"`
	// Description        string  `json:"description"`
	// AutoCreateApiToSql bool    `json:"autoCreateApiToSql"`
	// AutoMoveFile       bool    `json:"autoMoveFile"`
	Fields []Field `json:"fields"`
}

type Field struct {
	FieldName string `json:"fieldName"`
	// FieldDesc string `json:"fieldDesc"`
	FieldType string `json:"fieldType"`
	// FieldJson       string `json:"fieldJson"`
	// DataType        string `json:"dataType"`
	// DataTypeLong    string `json:"dataTypeLong"`
	// Comment         string `json:"comment"`
	// ColumnName      string `json:"columnName"`
	// FieldSearchType string `json:"fieldSearchType"`
	// DictType        string `json:"dictType"`
}
type Options struct {
	MailHost string
	MailPort int
	MailUser string // 发件人
	MailPass string // 发件人密码
	MailTo   string // 收件人 多个用,分割
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

func Send(o *Options) error {

	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", o.MailUser)

	//设置发送给多个用户
	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", o.Subject)

	//设置邮件正文
	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func main() {
	// Define a template.
	const letter = `
// 自动生成模板{{.StructName}}
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
		global.GVA_MODEL {{- range .Fields}}
			{{- if eq .FieldType "bool" }}
		{{.FieldName}}  *{{.FieldType}} 
			{{- else }}
		{{.FieldName}}  {{.FieldType}} 
			{{- end }} {{- end }}
}

{{ if .TableName }}
func ({{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
{{ end }}
`
	/*
		var recipients = []AutoCodeStruct{
			{"User", "tbl_User", []Field{
				Field{"username", "string"},
				Field{"userpwd", "string"}}}}
		// Create a new template and parse the letter into it.
		t := template.Must(template.New("letter").Parse(letter))
		// Execute the template for each recipient.
		for _, r := range recipients {
			err := t.Execute(os.Stdout, r)
			if err != nil {
				log.Println("executing template:", err)
			}
		}
	*/
	//连接字符串
	var connString string
	/*
		if u == "" || U == "" {
			connString = fmt.Sprintf("server=%s;port%d;trusted_connection=yes;database=%s", s, port, d)
		} else {
			connString = fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", s, port, d, u, p)
		}
	*/
	/*
		isdebug := true
		s := `172.17.3.144`
		port := 1433
		d := "master"
		var u, p string
		connString = fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", s, port, d, u, p)
		if isdebug {
			fmt.Println(connString)
		}
		//建立连接
		db, err := sql.Open("mssql", connString)
		err = db.Ping()
		if err != nil {
			log.Fatal("db ping err:", err)
			return
		}
		fmt.Println("err:", err)
		if err != nil {
			log.Fatal("Open Connection failed:", err.Error())
		}
		rows, err := db.Query(`select name from master..sysobjects where 1<>1`)
		defer rows.Close()
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				log.Fatal(err)
			}
			fmt.Println("name:", name)
		}
		defer db.Close()
	*/
	fmt.Println("conString:", connString)
	defer fmt.Println("a")
	defer fmt.Println("b")

	// currentDate := time.Now().Format("2006-01-02 15:04:05")
	// currentUnix := time.Now().Unix()
	// fmt.Println("currentDate:", currentDate)
	// fmt.Println("currentUnix:", currentUnix)

	// CurrentMilliUnix := time.Now().UnixNano() / 1000000
	// fmt.Println("CurrentMilliUnix:", CurrentMilliUnix)

	// CurrentNanoUnix := time.Now().UnixNano()
	// fmt.Println("CurrentNanoUnix:", CurrentNanoUnix)

	// uuid, _ := uuid.NewRandom()
	// uuid1 := uuid.String()
	// uuidtype := reflect.TypeOf(uuid1).String()
	// fmt.Println("uuid1:", uuid1)
	// fmt.Println("uuidtype:", uuidtype)

	// client, err := smtp.Dial("salarymail01.salarynet.local:25")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client.Mail("dba@salary.com")
	// client.Rcpt("ken.shi@salary.com")

	// wc, err := client.Data()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer wc.Close()
	// buf := bytes.NewBufferString("this is the email body.")
	// if _, err = buf.WriteTo(wc); err != nil {
	// 	log.Fatal(err)
	// }
	m := &Options{
		MailHost: "salarymail01.salarynet.local",
		MailPort: 25,
		MailUser: "dba@salary.com",
		MailTo:   "ken.shi@salary.com",
		Subject:  "Mail Test",
		Body:     "This is a test e-mail sent from Database Mail on SALARYSQL27.",
	}
	err := Send(m)
	if err != nil {
		panic(err)
	}
	fmt.Println("success!")

}
