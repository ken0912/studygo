package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

//LoadCompany ...
type LoadCompany struct {
	ID                int    `json:"id"`
	CompanyCode       string `json:"companycode"`
	CompanyName       string `json:"companyname"`
	CompanySize       string `json:"companysize"`
	ParentCompanyCode string `json:"parentcompanycode"`
	RunCode           string `json:"runcode"`
	RunDate           string `json:"rundate"`
	Segment           string `json:"segment"`
	FTEs              string `json:"ftes"`
	OwnershipCode     string `json:"ownershipcode"`
}

func main() {
	lc := LoadCompany{
		CompanyCode: "P110",
		CompanyName: "test",
		CompanySize: "100000",
	}
	fmt.Println("lc:", lc)

	lctype := reflect.TypeOf(lc)
	fmt.Println("lctype:", lctype)
	fmt.Println(lctype.NumField())
	for i := 0; i < lctype.NumField(); i++ {
		name := lctype.Field(i).Name
		path := lctype.Field(i).PkgPath
		vtype := lctype.Field(i).Type.Name()
		json := lctype.Field(i).Tag.Get("json")
		fmt.Println(name, path, vtype, json)

	}

	// lcvalue := reflect.ValueOf(lc)
	// fmt.Println("lcvalue:", lcvalue)

	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	// arr1 := arr
	// arr[1] = 10
	// fmt.Println("arr:", arr)
	// fmt.Println("arr1:", arr1)

	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := []int{5, 4, 3}

	// // copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	// // fmt.Println(slice1, slice2)
	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	// fmt.Println(slice1, slice2)

	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := []int{}
	// slice2 = append(slice2, slice1[:]...)
	// fmt.Println(&slice1[0], &slice2[0])

	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := slice1
	// fmt.Println(&slice1[0], &slice2[0])
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := i
		b := IsEven(a)
		// fmt.Println(rand.Intn(5))
		fmt.Println(a, b)
	}

	fmt.Println("--------------------------------------")

	for i := 1; i < 10000; i++ {
		a := math.Sqrt(float64(i) + 100)
		b := math.Sqrt(float64(i) + 168)
		if int(a)*int(a) == i+100 && int(b)*int(b) == i+168 {
			fmt.Println("i:", i)
			break
		}

	}

}
func IsEven(a int) bool {
	if a&1 == 0 {
		return true
	}
	return false
}
