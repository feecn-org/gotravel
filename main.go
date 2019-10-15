package main

import (
	"gotravel/cmd"
	"gotravel/custom"
)

func main() {
	//创建
	cmd.MakeDecion()
	custom.MakeCustomer()
	//travel.Mark()
	//println(prepare.GetReady())
	//println(tool.Http_Get("https://www.hao123.com"))
	//osCmd()
}

func osCmd() {
	cmd.OsCmd()
}
