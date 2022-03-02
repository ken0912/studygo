package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"gopkg.in/ini.v1"
)

type Host struct {
	Value []string `ini:"value,omitempty,allowshadow"`
}

type MssqlServiceInfo struct {
	ServerName   string        `json:"servername"`
	IPAddress    string        `json:"ipaddress"`
	ServicesInfo []ServiceInfo `json:"servicesinfo"`
}
type ServiceInfo struct {
	ServiceName        string `json:"servicename"`
	ServiceDisplayName string `json:"servicedisplayname"`
	ServiceStatus      string `json:"servicestatus"`
}

var host = new(Host)

func init() {
	cfg, err := ini.ShadowLoad("mssql_serverlist.ini")
	if err != nil {
		log.Fatal("ini.ShadowLoad() err:", err)
	}
	err = cfg.Section("Host").MapTo(host)
	if err != nil {
		log.Fatal("cfg.MapTo() err:", err)
	}

}

func main() {
	for _, v := range host.Value {
		msi := GetServicesFromServer(v)

		m, err := json.Marshal(msi)
		if err != nil {
			fmt.Println("json.Marshal err:", err)
		}
		ServicesData := string(m)
		fmt.Println("ServicesData:", ServicesData)
		// fmt.Println("mssqlserviceinfo:", msi.ServerName)
		// fmt.Println("mssqlserviceinfo:", msi.IPAddress)
		// fmt.Println("mssqlserviceinfo:", msi.ServicesInfo)
	}
}

func GetServicesFromServer(host string) MssqlServiceInfo {
	m, err := mgr.ConnectRemote(host)
	if err != nil {
		log.Fatal("connect err:", err)
	}
	defer m.Disconnect()

	ListServices, err := m.ListServices()
	if err != nil {
		log.Fatal("ListServices err:", err)
	}
	ms := new(MssqlServiceInfo)

	h, _ := net.LookupHost(host)
	ms.IPAddress = host
	ms.ServerName = h[0]

	for _, servicname := range ListServices {
		if strings.Contains(servicname, "SQL") {
			si := new(ServiceInfo)

			s, err := m.OpenService(servicname)
			if err != nil {
				log.Fatal("m.OpenService() err:", err)
			}
			status, err := s.Query()
			if err != nil {
				log.Fatal("Query() error:", err)
			}
			config, err := s.Config()
			if err != nil {
				log.Fatal("s.Config err:", s.Config)
			}
			si.ServiceName = servicname
			si.ServiceDisplayName = config.DisplayName
			si.ServiceStatus = GetStateDesc(status.State)
			ms.ServicesInfo = append(ms.ServicesInfo, *si)
		}

	}
	return *ms
	// if status.State == 1 {
	// 	fmt.Println("Starting the service:", service)
	// 	err = s.Start()
	// 	if err != nil {
	// 		log.Fatal("s.Start() err:", err)
	// 	}
	// 	fmt.Println(service, config.DisplayName, "StartPending")
	// }

	// //stop the service
	// status, err = s.Control(svc.Stop)
	// if err != nil {

	// 	fmt.Println("s.Control err:", err)
	// }
	// fmt.Println(service, config.DisplayName, GetStateDesc(status.State))

}

func GetStateDesc(s svc.State) string {
	switch s {
	case 1:
		return "Stoped"
	case 2:
		return "StartPending"
	case 3:
		return "StopPending"
	case 4:
		return "Running"
	case 5:
		return "ContinuePending"
	case 6:
		return "PausePending"
	case 7:
		return "Paused"
	}
	return ""
}
