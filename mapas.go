package datastructuresgo

import "fmt"

func Maps(){

	PrintMessage("\nMapas:")

	myMap := make(map[string]int)
		myMap["one"] = 1
		myMap["two"] = 2
	fmt.Println(myMap);
	fmt.Println("one:",myMap["one"]) // Output: 1
	
	value, ok := myMap["three"]
		if ok {
    		fmt.Println(value)
		} else {
    		fmt.Println("Key not found")
		}
	
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}
}