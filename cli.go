package main

import "fmt"
import "os"
//port "github.com/davecgh/go-spew/spew"	

func main() {
	component := os.Args[1]

	if component == "issue" {

		endpoint := "issue/GFG-7586"
		message := new(Message)
		doRequest(endpoint, message)		

		fmt.Println(message)
	}
}
