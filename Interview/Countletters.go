package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	JobTitle := `Retail Sales Consultant – Full Time/Part Time`
	re := regexp.MustCompile(`\b(?i)(FULL.?TIME|CONTRACTOR|INTERN|PART.?TIME|PER_DIEM|TEMPORARY|VOLUNTEER)\b`)
	value := re.FindAllString(JobTitle, -1)
	fmt.Println("value:", value)
	re = regexp.MustCompile(`[- ]`)
	for i, v := range value {
		value[i] = re.ReplaceAllString(v, "_")
		fmt.Println("newvalue:", value[i])
	}
	valstr := strings.Join(value, ",")
	fmt.Println("valstr:", valstr)
}
