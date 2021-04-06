package testcase02

import (
	"fmt"
	"testing"
)

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	err := monster.Store()
	if err != nil {
		t.Fatalf("monster.Store()错误:%v", err)
	}
	t.Logf("monster.Store()测试成功!")
}

func TestReStore(t *testing.T) {
	var monster Monster
	err := monster.ReStore()
	if err != nil {
		t.Fatalf("monster.Store()错误:%v", err)
	}
	fmt.Println("monster:", monster)
	t.Logf("monster.ReStore()测试成功!")
}
