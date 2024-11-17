package main

import "fmt"

func main() {
	fmt.Println("Maps")

	roles := make(map[string]string)

	roles["devops"] = "cloud"
	roles["java"] = "spring"
	roles["database"] = "db"

	fmt.Printf("type of the map %T\n", roles)
	fmt.Println("roles map here ", roles)

	// fetch the vaule of a particular key
	fmt.Println("fetch value ", roles["devops"])

	// add to the map
	fmt.Println("roles map here ", roles)

	// delete from map
	delete(roles, "database")
	fmt.Println("post delete here ", roles)

	// for loop with map
	for key, value := range roles {
		fmt.Printf("for key %v the value %v\n", key, value)
	}

	// key is not counted here
	// we put _ in the key part here
	// it is a format of comma, ok flow
	for _, value := range roles {
		fmt.Printf("the value %v\n", value)
	}

}
