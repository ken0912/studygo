package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// filepath := `D:\Project\Career.com\Data`

	// files, err := ioutil.ReadDir(filepath)
	// if err != nil {

	// 	log.Fatal(err)
	// }

	// for _, f := range files {
	// 	filepath := path.Join(filepath, f.Name())
	// 	log.Println("filepath:", filepath)
	// 	ReadText(filepath)
	// }
	ReadText("test")
}

func ReadText(fp string) {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(`D:\Project\Career.com\Data/JobPostingTitleToBenchmarkJob (2).txt`)

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	defer file.Close()

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string
	flag := 0
	for scanner.Scan() {
		text = append(text, scanner.Text())
		flag++

		if flag >= 10 {
			break
		}
	}

	// and then a loop iterates through
	// and prints each of the slice values.
	for _, each_ln := range text {
		fmt.Println(strings.Split(each_ln, "\t")[0])
	}
}
