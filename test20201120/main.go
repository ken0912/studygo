package main

import (
	"errors"
	"fmt"
)

//轮询负载均衡
type RoundRobinBalance struct {
	curIndex int
	rss      []string
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}

	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RoundRobinBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	lens := len(r.rss)
	if r.curIndex >= lens {
		r.curIndex = 0
	}

	curAddr := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return curAddr
}

func (r *RoundRobinBalance) Get() (string, error) {
	fmt.Println("r.curIndex:", r.curIndex)
	return r.Next(), nil
}

// "127.0.0.1:8001"
// "127.0.0.2:8001"
// "127.0.0.3:8001"
// "127.0.0.4:8001"
func main() {
	r := &RoundRobinBalance{}
	fmt.Println("r:", r)
	r.curIndex = 0
	r.rss = []string{
		"127.0.0.1:8001",
		"127.0.0.2:8001",
		"127.0.0.3:8001",
		"127.0.0.4:8001"}
	// err := r.Add(endpoint...)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	s, _ := r.Get()
	s1, _ := r.Get()
	s2, _ := r.Get()
	s3, _ := r.Get()
	fmt.Println("s:", s)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

}
