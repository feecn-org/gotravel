package main

import (
	"gotravel/cmd"
	"gotravel/prepare"
	"gotravel/tool"
	"gotravel/travel"
)

func main() {
	//创建
	systemCode := cmd.SystemCode()
	tool.InitLogger(systemCode)
	location := prepare.LoadSourceLocation()
	travel.Mark(location)
	//cmd.MakeDecion()
	//custom.MakeCustomer()
	//println(prepare.GetReady())
	//println(tool.Http_Get("https://www.hao123.com"))
	//osCmd()
}
