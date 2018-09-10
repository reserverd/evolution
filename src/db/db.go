package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

type StockData struct {
	Date        string
	Start       float32
	End         float32
	DiffPrice   float32
	DiffPrecent float32
	Low         float32
	High        float32
	Count       int
	Total       float32
	Change      float32
}

func InitDB(dbpath string) error {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}

	return nil
}

func Add(code string, sds []StockData) error {
	sql_table := `
    CREATE TABLE IF NOT EXISTS ` + code + ` (
        date DATE NULL,
        start REAL NULL,
	end REAL NULL,
	diffprice REAL NULL,
	diffprecent REAL NULL,
	low REAL NULL,
	high REAL NULL,
	count INTERGER NULL,
	total REAL NULL,
	change REAL NULL
    );
    `

	_, err = db.Exec(sql_table)
	if err != nil {
		return err
	}

	return nil
}

func ToStockData(stocks [][]string) (sds []StockData) {
	sds = make([]string, len(stocks))

	for i, stock := range stocks {
		sds[i].Data = stock[0]
		sds[i].Start = strconv.ParseFloat(stock[1])
		sds[i].End = strconv.ParseFloat(stock[2])
		sds[i].DiffPrice = strconv.ParseFloat(stock[3])
		sds[i].DiffPrecent = strconv.ParseFloat(stock[4])
		sds[i].Low = strconv.ParseFloat(stock[5])
		sds[i].High = strconv.ParseFloat(stock[6])
		sds[i].Total = strconv.ParseFloat(stock[7])
		sds[i].Change = strconv.ParseFloat(stock[8])
	}
}
