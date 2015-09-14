package main

import (
	"fmt"
	"github.com/mxk/go-sqlite/sqlite3"
)

func dbConnectionAndClose() {
	c, _ := sqlite3.Open("E:/go/dev/sample/src/sqlite/test.s3db")
	sql := "SELECT * FROM userinfo"
	row := make(sqlite3.RowMap)
	for s, err := c.Query(sql); err == nil; err = s.Next() {
		var rowid int64
		s.Scan(&rowid, row)     // Assigns 1st column to rowid, the rest to row
		fmt.Println(rowid, row) // Prints "1 map[a:1 b:demo c:<nil>]"
	}

	defer c.Close()
}

func main() {
	dbConnectionAndClose()
	fmt.Println("SUCCESS!")
}
