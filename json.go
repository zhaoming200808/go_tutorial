package main

import "fmt"
import "os"
import "encoding/json"

func main() {
	test_fun()
	print()
	t2()

}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func test_fun() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
	return
}

func t2() {
	var jsonBlob = []byte(`[
			{"Name": "Platypus", "Order": "Monotremata"},
			{"Name": "Quoll",    "Order": "Dasyuromorphia"}
			]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
