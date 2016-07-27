package main

import "os"
import "fmt"

type Issue struct {
	Fields struct {
		Assignee struct {
			Self string
			Name string
			DisplayName string
		}
		Reporter struct {
			Name string
			DisplayName string

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
		Summary string
		CreatedAt string `json:"created"`
	}	
}

func (i *Issue) show() {
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Summary      : %s\n", i.Fields.Summary)
	fmt.Printf("Reporter     : %s (user: %s)\n", i.Fields.Reporter.DisplayName, i.Fields.Reporter.Name)
	fmt.Printf("Assinged to  : %s (user: %s)\n", i.Fields.Assignee.DisplayName, i.Fields.Assignee.Name)
	var fixVersions string;
	for _, version := range i.Fields.FixVersions {
		fixVersions += version.Name + ","
	}
	fmt.Printf("Fix Versions : %s\n", fixVersions)
	var components string;
	for _, comp := range i.Fields.Components {
		components += comp.Name + ","
	}
	fmt.Printf("Components   : %s\n", components)
	fmt.Printf("Status       : %s\n", i.Fields.Status.Name)
	fmt.Printf("Created At   : %s\n", i.Fields.CreatedAt)
	fmt.Println("-------------------------------------------------")
}

func RunIssue() {
	action := os.Args[2]
	if action == "get" {
		id := os.Args[3]
		//endpoint := "issue/GFG-7586"
		endpoint := "issue/" + id
		issue := new(Issue)
		doRequest(endpoint, issue)

		issue.show()
	}
}