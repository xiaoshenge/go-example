package main

import "fmt"

func keys(m map[interface{}]interface{}) []interface{} {
	keys := make([]interface{}, 0)

	for k, _ := range m{
		keys = append(keys, k)
	}
	return keys
}
func main() {
	m := make(map[string]string, 1)
	m["demo"] = "data"
	fmt.Printf(keys(m))
}