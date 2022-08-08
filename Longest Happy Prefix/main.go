package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestPrefix("bba"))
}

func longestPrefix(s string) string {
	var res string
	// loop through every prefix and see if it's a suffix
	for i := 1; i < len(s); i++ {
		if strings.HasSuffix(s, s[0:i]) {
			res = s[0:i]
		}
	}

	return res
}
