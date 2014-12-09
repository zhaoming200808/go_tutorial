package main

import "fmt"

type Tree struct {
	Success bool `json:"success"`
	flag    string
	Data    []interface{} `json:"data"`
}

type Node struct {
	Id       int           `json:"id"`
	Pid      int           `json:"pid"` //if ( Id ==  Pid ) is Root
	Name     string        `json:"name"`
	Children []interface{} `json:"children"`
}

func init() {
	//	var roots []Node
	//	var leafs []Node
	//	var nodes []Node
}

func main() {
	var trees []Tree
	//	var tree Tree
	var itrees []interface{}

	fmt.Printf("=============================================")
	tree_1 := Tree{Success: true, flag: "tree_1"}
	tree_2 := Tree{Success: true, flag: "tree_2"}
	trees = append(trees, tree_1, tree_2)
	for _, tree := range trees {
		itrees = append(itrees, tree)
	}

	//	tree = Tree{Success: true, flag: "tree", Data: itrees}

	fmt.Printf("=============================================")
}
