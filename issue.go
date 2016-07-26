package main

import "os"
import "fmt"

type Message struct {
	Fields struct {
		Assignee struct {
			Self string `json:"self"`
			Name string `json:"Name"`
		}
		Reporter struct {
			Name string

		}
		Components []struct {
			Name string
			Description string
		}
		FixVersions []struct {
			Description string
			Name string
		}
		Status struct {
			Name string
		}
	}	
}

type Issue struct {

}
func (i *Issue) Run() {
	params := os.Args[2:]
	fmt.Println(params)
}