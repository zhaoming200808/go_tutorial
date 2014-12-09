package main

import "os"
import "fmt"
import "encoding/json"

func main() {
	var d []interface{}
	leaf_1 := Tree{Id: 10001, Text: "leaf_1"}
	leaf_2 := Tree{Id: 10002, Text: "leaf_2"}
	var l []Tree
	l = append(l, leaf_1, leaf_2)
	tree_1 := Tree{Id: 1, Text: "tree_1", Children: l}
	d = append(d, tree_1)
	menuTree := &MenuTree{Success: true, Date: d}
	fmt.Printf("%#v", menuTree)
	b, err := json.Marshal(menuTree)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print("\n")
	os.Stdout.Write(b)
}

type MenuTree struct {
	Success bool          `json:"success"`
	Date    []interface{} `json:"date"`
}

type Tree struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Leaf     bool   `json:"leaf"`
	Children []Tree `json:"children"`
}
