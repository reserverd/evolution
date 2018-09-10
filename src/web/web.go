package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStockData(code, start, end string) (string, error) {
	resp, err := http.Get("http://q.stock.sohu.com/hisHq?code=" + code + "&start=" + start + "&end=" + end + "&stat=1&order=D&period=d&callback=historySearchHandler&rt=jsonp")
	if err != nil {
		fmt.Printf("http get failed: %v\n", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read all from resp.body failed: %v\n", err)
		return "", err
	}

	return string(body), nil
}
