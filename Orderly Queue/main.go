package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(orderlyQueue("gxzv", 4))
}

func orderlyQueue(s string, k int) string {
	ans := s
	// if k grater than 1 then we just sort it, be cause we can create
	// any combo we want with enogh changes
	if k != 1 {
		s1 := strings.Split(s, "")
		sort.Strings(s1)
		return strings.Join(s1, "")
	}
	// if k == 1 then to go through all the possible options
	// and find the min
	for i := 0; i < len(s); i++ {
		tmp := fmt.Sprintf("%s%c", s[1:], s[0])
		if tmp < ans {
			ans = tmp
		}
		// updating the string we want to check
		s = tmp
	}

	return ans
}
