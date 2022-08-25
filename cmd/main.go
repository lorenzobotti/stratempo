package main

import (
	"flag"
	"fmt"
	"os"
	"stratempo"
	"strings"
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "give me a number")
		os.Exit(1)
	}

	arg := strings.Join(os.Args[1:], " ")
	converted, err := stratempo.Convert(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// converted, _ := stratempo.Convert(" ")

	fmt.Println(converted)
}
