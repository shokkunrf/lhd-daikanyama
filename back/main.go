package main

import "lhd-daikanyama/router"
import _ "github.com/go-sql-driver/mysql"

func main() {
	r, err := router.GetRoute()
	if err != nil {
		panic(err)
	}
	r.Run()
}
