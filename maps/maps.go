package main

import "fmt"

func main() {
	userOne := map[string]string{
		"name":     "Max",
		"lastname": "Rodrigues",
	}
	// fmt.Println(user.name) Not valid
	fmt.Println(userOne["name"])

	userTwo := map[string]map[string]string{
		"name": {
			"first": "Max",
			"last":  "Rodrigues",
		},
		"course": {
			"name": "Computer Cience",
		},
	}

	fmt.Println(userTwo)
	delete(userTwo, "course") // delete map key
	fmt.Println(userTwo)

	//add new key
	userTwo["mom"] = map[string]string{
		"name": "Eurides",
	}

	fmt.Println(userTwo)
}
