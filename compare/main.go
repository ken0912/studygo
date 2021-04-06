package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	apath := `D:\Test\CAE\Ecomp_Salary`
	bpath := `D:\Test\CAE\Ecomp_Template`
	PthSep := string(os.PathSeparator)
	flist, err := ioutil.ReadDir(apath)
	if err != nil {
		panic(err)
	}
	n := 0
	for _, v := range flist {
		filename := v.Name()
		afullpath := apath + PthSep + filename
		bfullpath := bpath + PthSep + filename
		a, _ := ioutil.ReadFile(afullpath)
		b, _ := ioutil.ReadFile(bfullpath)
		acontent := string(a)
		bcontent := string(b)
		acontent = strings.Replace(acontent, "\n", "", -1)
		bcontent = strings.Replace(bcontent, "\n", "", -1)
		if string(acontent) == string(bcontent) {
			continue
		} else {
			n += 1
			fmt.Println("filename:", filename)
		}

	}

}
