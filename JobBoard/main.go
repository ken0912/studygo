package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Stat(`D:\Work\TFS\Compensation\Compensation\Consumer\Database\SalaryCommunity\DDL\usp_SCMU_GetUserDiscussionByUserID.sql`)
	fmt.Println("f:", f)
	fmt.Println("err:", err)
	fmt.Println(os.IsNotExist(err))
}
