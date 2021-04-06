package main

// import "time"
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
func main()  {
	

	sourcepath := `D:\Test\20200624`
	targetpath := `D:\Test\2020062401`
	l,err:=ioutil.ReadDir(sourcepath)
	if err!=nil{
		panic(err)
	}
	var filename string
	for _,v:=range l{
		filename = v.Name()
		if strings.HasSuffix(filename,".sql"){
			fmt.Println("filename:",filename)

			// fmt.Println(targetpath+"\\"+filename)
			fullpath :=sourcepath+"\\"+filename
			content,err := ioutil.ReadFile(fullpath)
			if err!=nil {
				fmt.Println("err:",err)
				panic(err)
			}
			newcontent:=strings.Replace(string(content),"ExecutiveCompensation..","ExecutiveCompensation_FreeTrial..",-1)
			
			f,err:=os.OpenFile(targetpath+"\\"+filename,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0644)
			if err!=nil{
				fmt.Println("err:",err)
				panic(err)
			}
			defer f.Close()
			n,err:=f.WriteString(newcontent)
			if err!=nil{
				fmt.Println("err:",err)
				panic(err)
			}
			// err = ioutil.WriteFile(targetpath+"\\"+filename,[]byte(newcontent),os.ModeAppend)
			if  err!=nil {
				fmt.Println("err:",err)
				panic(err)
			}
			fmt.Printf("%v 成功写入%v \n",filename,n)
		}
		
		// fmt.Printf("[%v]: %v:",filename,newcontent)
	}
}

