package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	//x := 1
	//println(x)
	//p := &x
	//println(p)
	//println(&x)
	//
	//println(*p)
	//*p = 2
	//println(*p)
	//println(x)

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
