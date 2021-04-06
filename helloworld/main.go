package main

import (
	"flag"
	"fmt"
)

func main() {
	//九九乘法表
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d x %d = %d\t", j, i, j*i)
	// 	}
	// 	fmt.Printf("\n")
	// }
	// for i := 9; i >= 1; i-- {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d x %d = %d	", j, i, j*i)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// s := "hello"

	// for _, v := range s {
	// 	fmt.Println(string(v))
	// }
	// scanf()
	// scan()
	// scanln()
	// s1 := "hello"
	// fmt.Println(string(s1[1]))
	// sArray()
	// a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// a2 := a1[5:10]
	// fmt.Println(a2)
	// sliceAppend()

	// var m1 map[string]int
	// m2 := []string{}
	// m3 := []string{"a", "b", "c"}
	// m1 = make(map[string]int, 10)
	// fmt.Printf("%p", m1)
	// fmt.Printf("%p", m2)
	// m1["'xiaoming"] = 98
	// fmt.Println(m2)
	// m2 = append(m2, m3...)
	// fmt.Println(m2)

	// var count int
	// s1 := "hello 你好"
	// for _, c := range s1 {
	// 	if unicode.Is(unicode.Han, c) {
	// 		count++
	// 	}
	// }
	// fmt.Println("汉字个数：", count)

	// s1 := "how do you do do"
	// s2 := strings.Split(s1, " ")
	// m1 := make(map[string]int, 10)

	// for _, w := range s2 {
	// 	m1[w]++
	// }
	// if i, ok := m1["do"]; ok {
	// 	fmt.Println("i", i)
	// 	fmt.Println("ok", ok)
	// }

	// sArray()

	// a := "hello world你好世界"
	// for i, v := range a {
	// 	fmt.Println("i--v", i, string(v))
	// }

	// var buf [16]byte
	// os.Stdin.Read(buf[:])
	// os.Stdout.WriteString(string(buf[:]))

	// var str string
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Please input:")
	// str, err := reader.ReadString('\t')
	// if err == nil {
	// 	fmt.Println(str)
	// }

	// fmt.Printf("%T", os.Args)
	// fmt.Println("args[0]:", os.Args[0])
	// if len(os.Args) > 1 {
	// 	for index, v := range os.Args {
	// 		if index == 0 {
	// 			continue
	// 		}
	// 		fmt.Printf("args[%d]: %s\n", index, v)
	// 	}
	// }

	var aoo string
	flag.StringVar(&aoo, "a", "A", "a is name")
	flag.Parse()

	fmt.Println(aoo)

}

func scanf() {
	var a string
	fmt.Scanf("%s", &a)
	fmt.Println("你输入的是:", a)
}

func scan() {
	var b string
	var c string
	fmt.Scan(&b, &c)
	fmt.Println("this is input:", b, c)
}

func scanln() {
	var b string
	var c string
	fmt.Scanln(&b, &c)
	fmt.Println("this is input:", b, c)
}

func sArray() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0
	for i, v := range a1 {
		sum += v
		fmt.Println("i--v", i, v)
	}
	fmt.Println("sum:", sum)
	fmt.Printf("%T", a1)
}

func sliceAppend() {
	s1 := []string{}
	s2 := []string{"beijing", "shanghai", "guangzhou"}
	s1 = append(s1, s2...)
	fmt.Println(s1)

}
