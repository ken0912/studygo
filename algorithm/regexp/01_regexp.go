package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	cleanCompanyName := GetDeepCleanCompanyName("153 salary.com and, Inc. (US)")
	fmt.Println("cleanCompanyName:", cleanCompanyName)

	cleanCompanyName = GetDeepCleanCompanyName("Kellogg Brown & Root LLC")
	fmt.Println("cleanCompanyName:", cleanCompanyName)

	cleanCompanyName = GetDeepCleanCompanyName("Snyder's-Lance, Inc.")
	fmt.Println("cleanCompanyName:", cleanCompanyName)

	re := regexp.MustCompile(`\w+`)
	matchedList := re.FindAllString(`salary.com`, -1)
	fmt.Println("matchedList", matchedList)
}

func GetDeepCleanCompanyName(companyName string) (cleanCompanyName string) {
	if companyName == "" {
		return
	}
	// 153 salary.com, Inc. (US)   to  153 salary.com, Inc.
	re := regexp.MustCompile(`\(.+\)`)
	cleanCompanyName = re.ReplaceAllString(companyName, "")
	// 153 salary.com, Inc. to salary.com
	re = regexp.MustCompile(`\w+.com\b`)
	matchWeSitevalue := re.FindString(companyName)
	if matchWeSitevalue != "" {
		// fmt.Println("matchWeSitevalue:", matchWeSitevalue)
		return matchWeSitevalue
	}
	// Kellogg Brown & Root LLC to Kellogg Brown Root; Snyder's-Lance, Inc. to Snyder Lance
	re = regexp.MustCompile(`\b(?i)(s|and|of|in|at|by|the|for|incorporated|corporation|Corporate|inc|llc|corp|group|company|limited|co|LLP|LP|GP|L.L.C)\b`)
	cleanCompanyName = re.ReplaceAllString(cleanCompanyName, "")
	re = regexp.MustCompile(`\w+`)
	matchedValue := re.FindAllString(cleanCompanyName, -1)

	cleanCompanyName = strings.Join(matchedValue, " ")

	return
}
