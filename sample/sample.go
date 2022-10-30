package sample

import "fmt"

func Sample() bool {
	var arr1 []string = []string{"hoge", "fuga", "bann"}
	var arr2 []string = []string{"neko", "fuga", "bann"}

	for i, v := range arr1 {
		if v != arr2[i] {
			fmt.Printf("false index: %v\n", i)
			return false
		}
	}
	return false
}
