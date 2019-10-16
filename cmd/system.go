package cmd

import (
	//"os"
	"fmt"
	"runtime"
)

func SystemCode() string {
	systemCode := runtime.GOOS
	fmt.Println(systemCode)
	return systemCode
}
