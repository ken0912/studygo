package main

import (
	"database/sql"
	"fmt"
	"os"
	"scandonetypefile/logger"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/kardianos/service"
)

var log = logger.NewFileLog("debug", "G:\\Ken\\", "log.log", 1024*10*1024)
var checkfile = `\\10.100.10.51\g$\EnterpriseSQL\IPAS\GSYBcpDataStoreForDailyUpdate\UploadToFTP.done`
var donefile = `\\10.100.10.51\g$\EnterpriseSQL\IPAS\GSYBcpDataStoreForDailyUpdate\DownloadAndUpdate.done`
var errfile = `\\10.100.10.51\g$\EnterpriseSQL\IPAS\GSYBcpDataStoreForDailyUpdate\DownloadAndUpdate.err`

type program struct {
	service service.Service
}

func (p *program) Start(s service.Service) error {
	p.service = s
	go p.run()
	return nil
}

func (p *program) run() {

	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		err := Exists(checkfile)

		if err {
			err := os.Remove(checkfile)
			if err != nil {
				log.Error("Remove error:%v", err)
				return
			}

			err = Process()
			if err != nil {
				f, _ := os.Create(errfile)
				f.Close()
				log.Error("err:%v", err)
				continue
			} else {
				f, _ := os.Create(donefile)
				f.Close()
				log.Trace("Downlaod and update tables success!")
				continue
			}

		}

	}

}

func (p *program) Stop(s service.Service) error {
	return nil
}

/**
* MAIN函数，程序入口
 */

func main() {

	svcConfig := &service.Config{
		Name:        "DBA_DownloadAndUpdateTables",       //服务显示名称
		DisplayName: "DBA Download and update tables",    //服务名称
		Description: "DBA Download and update tables...", //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal("%v", err)
	}

	// logger 用于记录系统日志

	if err != nil {
		log.Fatal("%v", err)
	}
	if len(os.Args) == 2 { //如果有命令则执行
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal("%v", err)
		}
	} else { //否则说明是方法启动了
		err = s.Run()
		if err != nil {
			log.Error("%v", err)
		}
	}
	if err != nil {
		log.Error("%v", err)
	}

}
func Exists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
func Process() error {
	var (
		s    = "SE-SQL-110"
		port = 1433
		d    = "IPAS_DBForUpdate"
	)
	connString := fmt.Sprintf("server=%s;port%d;trusted_connection=yes;database=%s", s, port, d)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:%v", err.Error())
		return err
	}
	defer db.Close()
	sqlstr := "exec usp_SyncDataFromStag"
	_, err = db.Exec(sqlstr)
	if err != nil {
		log.Fatal("Exec usp_SyncDataFromStag failed:%v", err.Error())
		return err
	}
	return nil
}
