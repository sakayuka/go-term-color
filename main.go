package main

import (
	"fmt"
)

var (
	fg = "\x1b[38;5;"
	bg = "\x1b[48;5;"
	rs = "\x1b[0m"
)

func main() {
	for i := 0; i < 256; i++ {
		fmt.Print(color(i))
		if (i+1)%8 == 0 {
			fmt.Print("\n")
		}
	}
}

func color(code int) string {
	number := fmt.Sprintf("%3d", code)
	return fmt.Sprintf("%s%dm %s%s%s%dm %s%s", bg, code, number, rs, fg, code, number, rs)
}
