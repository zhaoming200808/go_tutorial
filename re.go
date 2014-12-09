package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("ab", "seafood")
	fmt.Println(matched, err)
}
