package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename) //return slice
		if err != nil {
			if err != nil {
				fmt.Println(err)
			}
		}
		for _, lien := range strings.Split(string(data), "\n") {
			counts[lien]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
	}
}
