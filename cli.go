package main

import "os"
//port "github.com/davecgh/go-spew/spew"	

func main() {
	component := os.Args[1]

	if component == "issue" {
		RunIssue()
	}
}
