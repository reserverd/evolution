package main

import (
	"db"
	"encoding/json"
	"fmt"
	"strings"
	"web"
)

type webresp struct {
	Status int `json:"status"`
	Hq     [][]string
	Code   string `json:"code"`
	Stat   []interface{}
}

func main() {
	var resp webresp

	if err := db.InitDB("/root/stock.db"); err != nil {
		fmt.Printf("Init DB failed: %v\n", err)
		return
	}

	// historySearchHandler([{"status":0,"hq":[["2018-08-17","56.90","56.67","0.09","0.16%","56.07","58.10","52051","29815.98","0.48%"],["2018-08-16","57.77","56.58","-1.45","-2.50%","56.57","58.26","70538","40538.68","0.65%"]],"code":"cn_600009","stat":["累计:","2018-08-16至2018-08-17","-1.36","-2.34%",56.07,58.26,122589,70354.66,"1.13%"]}])
	data, err := web.GetStockData("cn_600009", "20180816", "20180817")
	if err != nil {
		fmt.Println(err)
		return
	}

	if !strings.HasPrefix(data, "historySearchHandler") {
		fmt.Printf("fetch web data error: %v", data)
		return
	}

	data = strings.TrimPrefix(data, "historySearchHandler([")
	data = strings.TrimSuffix(data, "])\n")

	err = json.Unmarshal([]byte(data), &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("web=%+v\n", resp)

	return
}
