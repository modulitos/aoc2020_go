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

	// O(N lg N)
	for i, a := range num {
		if a > 1010 {
			// short-circuit, since the next number would
			// have to be greater than 1010, resulting in
			// a sum greater than 2020.
			break
		}

		if a == 1010 {
			if i+1 < len(num) && num[i+1] == 1010 {
				fmt.Println("answer part 1:", 1010, 1010, 1010+1010)
			}
			continue
		}
		b := 2020 - a
		// sort.Search does a binary search. Returns len(num)
		// if no match is found.
		// https://golang.org/pkg/sort/#Search
		j := sort.Search(len(num), func(i int) bool { return num[i] >= b })
		if j < len(num) && num[j] == b {
			fmt.Println("answer part 1:", a, b, a*b)
		}
	}

	// O(N^2 lg N)
	for i, a := range num {
		if a > (2010/3)+2 {
			break
		}
		for j, b := range num[i+1:] {
			c := 2020 - a - b
			slice := num[j:len(num)]
			k := sort.Search(len(slice), func(i int) bool { return num[i] >= c })

			if k < len(slice) && slice[k] == c {
				fmt.Println("answer part 2:", a, b, c, a+b+c)
			}
		}
	}
}
