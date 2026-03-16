package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "2.0.0"

func greeting(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	name := flag.String("name", "", "Name to greet")
	showVersion := flag.Bool("version", false, "Show version")
	flag.BoolVar(showVersion, "v", false, "Show version (short)")
	noColor := flag.Bool("no-color", false, "Disable color output")
	flag.Parse()

	if *showVersion {
		fmt.Printf("hello-world v%s\n", version)
		os.Exit(0)
	}

	msg := greeting(*name)
	if *noColor || os.Getenv("NO_COLOR") != "" {
		fmt.Println(msg)
	} else {
		fmt.Printf("\033[32m%s\033[0m\n", msg)
	}
}
