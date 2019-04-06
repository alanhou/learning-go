package main

import "fmt"

func main() {
	m := map[string] string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string] int) // m2 == empty map

	var m3 map[string] int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map:")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values:")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok := m["cause"]; ok{ // 不存在的 key 为 Zero value
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Delete values:")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
