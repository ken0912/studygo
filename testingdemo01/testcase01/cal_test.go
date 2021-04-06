package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10) 执行错误,期望值=%v,实际值=%v", 55, res)
	}
	t.Logf("AddUpper(10) 执行正确...")
}
