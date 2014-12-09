package main

//import "os"
import (
	"fmt"
	"strconv"
)

func main() {
	sl := []int{}
	sl = append(sl, 1, 2, 3, 4, 37, 36, 38, 40, 41, 42, 43, 44, 45, 46, 77, 78, 113)
	fmt.Printf("sl: |%v|\n", sl)
	name := getDBName(sl)
	fmt.Printf("name: |%s|\n", name)
}

func getDBName(sl []int) (name string) {
	// sort
	count := len(sl)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}

	// create name
	if count < 2 {
		return
	} else if count == 2 {
		name = "Z" + strconv.Itoa(sl[0]) + "_Z" + strconv.Itoa(sl[1])
		return name
	}

	flag := false
	for i := 0; i < count; i++ {
		num := strconv.Itoa(sl[i])

		if flag == true {
			if i+1 < count {
				if sl[i]+1 != sl[i+1] {
					name = name + "Z" + strconv.Itoa(sl[i])
					flag = false
				}
			}
			if i+1 == count {
				name = name + "Z" + strconv.Itoa(sl[i])
			}
		} else {

			if i+1 < count && sl[i]+1 == sl[i+1] {
				if sl[i]+2 == sl[i+2] {
					flag = true
				}
			}

			if name == "" {
				name = name + "Z" + num
			} else {
				name = name + "_Z" + num
			}

		}

	}
	return name
}

func test_fun() {
	fmt.Printf("i am test func \n")
}
