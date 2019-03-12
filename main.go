package main

import (
	"gotravel/prepare"
	"gotravel/travel"
)

func main() {
	travel.Mark()
	println(prepare.GetReady())
}
