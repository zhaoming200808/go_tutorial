package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	var re_num = regexp.MustCompile(`^[1-9][0-9]+$`)

	fmt.Println("100", re_num.MatchString("100"))
	fmt.Println("10a", re_num.MatchString("10a"))
	fmt.Println("a10", re_num.MatchString("a10"))
	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
}

