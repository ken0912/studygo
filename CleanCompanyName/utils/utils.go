package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/demdxx/gocast"
)

func GetDeepCleanCompanyName(companyName string) string {
	if companyName == "" {
		return companyName
	}
	// 153 salary.com, Inc. (US)   to  153 salary.com, Inc.
	re := regexp.MustCompile(`\(.+\)`)
	cleanCompanyName := re.ReplaceAllString(companyName, "")
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

	if cleanCompanyName == "" {
		return companyName
	}
	return cleanCompanyName
}

func SQLToJSON(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			// var v interface{}
			val := values[i]

			entry[col] = convertRow(val)
		}
		// fmt.Println("entry:", entry)
		tableData = append(tableData, entry)
	}
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(tableData)
	if err != nil {
		panic(err)
	}

	return byteBuf.String(), nil
}

func convertRow(row interface{}) interface{} {
	switch row.(type) {
	case int:
		return gocast.ToInt(row)
	case string:
		return gocast.ToString(row)
	case []byte:
		return gocast.ToString(row)
	case bool:
		return gocast.ToBool(row)
	case float32:
		return gocast.ToFloat(row)
	case float64:
		return gocast.ToFloat(row)
	}
	return row
}
