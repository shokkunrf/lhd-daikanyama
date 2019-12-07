package main

import "lhd-daikanyama/router"

func main() {
	r := router.GetRoute()
	r.Run()
}
