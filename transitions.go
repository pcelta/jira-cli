package main

import "os"
import "fmt"

type TransitionList struct {
	Transitions []struct {
		Id string
		Name string
		To struct {
			Self string
			Name string
			Id string		
		}
	}
}

func (t *TransitionList) show() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("Possible Transitions----------------------------")
	fmt.Println("-------------------------------------------------")
	fmt.Println("-------------------------------------------------")
	for _, transit := range t.Transitions {
		fmt.Printf("Status Id     : %s\n", transit.To.Id)
		fmt.Printf("Status Name   : %s\n", transit.To.Name)
		fmt.Println("-------------------------------------------------")
		fmt.Println("-------------------------------------------------")
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("Use: > transition move <TICKET-ID> <STATUS-ID>---")
}

func RunTransitions() {
	action := os.Args[2]
	if action == "get" {
		id := os.Args[3]
		//endpoint := "issue/GFG-7586"
		endpoint := fmt.Sprintf("issue/%s/transitions", id)
		transitions := new(TransitionList)
		doRequest(endpoint, transitions)

		transitions.show()
	}
}