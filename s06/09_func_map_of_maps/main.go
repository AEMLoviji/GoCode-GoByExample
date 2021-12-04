package main

import "fmt"

func main() {
	employees := map[string]map[string]string{
		//also possible to write as below, but it is redundant
		//"BT2": map[string]string{
		//	"firstName": "Blake",
		//	"lastName":  "Travis",
		//},
		"BT": {
			"firstName": "Blake",
			"lastName":  "Travis",
		},
		"PC": {
			"firstName": "Parker",
			"lastName":  "Cooper",
		},
		"DC": {
			"firstName": "Dakota",
			"lastName":  "Carrington",
		},
	}

	if emp, ok := employees["PC"]; ok {
		fmt.Println(emp["firstName"], emp["lastName"])
	}

	for initials, emp := range employees {
		fmt.Println(initials, emp["firstName"], emp["lastName"])
	}
}
