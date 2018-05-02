package main

import (
	"fmt"
	"path"
)

func helpBaseUsage(arguments []string) {
	//fmt.Println("Help message for the following CLI command:")
	fmt.Println(path.Base(arguments[0]))
	fmt.Println("github - GitHub command line interface")
	fmt.Println("\nThis is an interface for your GitHub account and your projects, repos and issues")
	return
}