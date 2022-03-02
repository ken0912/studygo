package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"orderserver/routers"
	"time"

	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/micro/v3/client/selector"
	"github.com/micro/micro/v3/registry"
	"github.com/micro/micro/v3/web"
)

var consulReg registry.Registry

func init() {
	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
	consulReg = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
}

func main() {
	//初始化路由
	ginRouter := routers.InitRouters()

	//注册服务
	microService := web.NewService(
		web.Name("orderserver"),
		//web.RegisterTTL(time.Second*30),//设置注册服务的过期时间
		//web.RegisterInterval(time.Second*20),//设置间隔多久再次注册服务
		web.Address(":18002"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)

	//获取服务地址
	hostAddress := GetServiceAddr("userserver")
	fmt.Println("hostAddress:", hostAddress)
	if len(hostAddress) <= 0 {
		fmt.Println("hostAddress is null")
	} else {
		url := "http://" + hostAddress + "/users"
		response, _ := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte("")))
		resData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("ioutil ready error:", err)
		}

		defer response.Body.Close()
		fmt.Println("resData:", string(resData))
	}

	microService.Run()
}

func GetServiceAddr(serviceName string) (address string) {
	var retryCount int
	for {
		servers, err := consulReg.GetService(serviceName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var services []*registry.Service
		for _, value := range servers {
			fmt.Println(value.Name, ":", value.Version)
			services = append(services, value)
		}
		next := selector.RoundRobin(services)
		if node, err := next(); err == nil {
			address = node.Address
		}
		if len(address) > 0 {
			return
		}
		//重试次数++
		retryCount++
		time.Sleep(time.Second * 1)
		//重试5次为获取返回空
		if retryCount >= 5 {
			return
		}
	}
}
