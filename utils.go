package gutils

import "fmt"

func Printfln(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Println()
}
