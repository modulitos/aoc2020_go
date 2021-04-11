package main

import (
	"fmt"
	"io"
	"log"
	"sort"
)

func main() {
	var num []int
	for {
		var x int
		_, err := fmt.Scanf("%d", &x)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		num = append(num, x)
	}

	sort.Ints(num)
	fmt.Println("nums, sorted:", num)
}
