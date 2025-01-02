package main

import "fmt"

func printx() string {

	var res string

	for i := 0; i < 10; i++ {
		res += fmt.Sprintf("MIC testing %d \n", i)
	}
	return res
}
