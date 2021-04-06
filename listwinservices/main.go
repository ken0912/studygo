package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func isProcessExist(appName string) (bool, string, int) {
	appary := make(map[string]int)
	cmd := exec.Command("cmd", "/C", "tasklist")
	output, _ := cmd.Output()
	//fmt.Printf("fields: %v\n", output)
	n := strings.Index(string(output), "System")
	if n == -1 {
		fmt.Println("no find")
		os.Exit(1)
	}

	data1 := string(output)
	fmt.Println("data1:", data1)
	data := string(output)[n:]
	fields := strings.Fields(data)
	for k, v := range fields {
		if v == appName {
			appary[appName], _ = strconv.Atoi(fields[k+1])

			return true, appName, appary[appName]
		}
	}

	return false, appName, -1
}

func main() {
	// var test = "asdfs Hello world sdf home officehome mortgage sta-home"
	// r, _ := regexp.Compile(`\bHello world\b`)
	// fmt.Println(r.FindAllString(test, -1))
	// fmt.Println(len(r.FindAllString(test, -1)))
	var hostname string
	ip := "salarysql14"
	domainsuffix := ".salarynet.local"
	addrs, err := net.LookupHost(ip + domainsuffix)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		hostname = strings.Join(addrs, " ")
	}
	fmt.Println("hostname:", hostname)

}
