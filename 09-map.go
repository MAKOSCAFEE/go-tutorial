package main

import "fmt"

/*
 * Hash tables - Contains key/value pair, key has to be unique. helps to access large pool of data.
 * Hash function is defined and is used to take a key and compute a slot in the hash table to insert the value according to the key.
 * A hash function is a function that takes as its argument the key and it returns the slot where you want to put the value.
 * A map - is the Go implementation of hashtable. Use make() to create map.
 * Accessing the value in the map youuse [key]. if not present it returns zero.
 */

func main() {
	// here key type is string and value type is int.
	// this is declations and making a variable.
	var idMap map[string]int

	// Assign the variable idMap, pointing to a map, making an actual map.
	// this will create an empty map.
	idMap = make(map[string]int)

	idMap["joe"] = 20

	// map literal
	idMap2 := map[string]int{"joe": 45}

	// map accessing
	fmt.Println(idMap2["joe"]) // 20

	// adding a key/value pair
	idMap2["jane"] = 50

	// deleting a key/value pair

	delete(idMap, "joe")

	// Two Value assignement,
	// in this case id is value, p is presence of the key(will return true if key is in the map)

	id, p := idMap2["jane"]

	// function len() to get to know the length of the map.

	fmt.Println(id, p, len(idMap2))

	// Iteration throughh a map

	for key, val := range idMap2 {
		fmt.Println(key, val)
	}

}
