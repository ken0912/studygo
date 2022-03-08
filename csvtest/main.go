package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open(`D:\Test\Legal_National_Detail_Annual_Oct_21.csv`)
	if err != nil {
		log.Fatalln("os.Open error:", err)
	}

	defer f.Close()
	r := csv.NewReader(f)
	r.LazyQuotes = true
	id := 0

	for {
		id++
		record, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}

		fmt.Println("record:", record)

		// recorditf := make([]interface{}, columnlen)
		// for i, v := range record {
		// 	recorditf[i] = v
		// }
		if id >= 10 {
			break
		}

	}
}
