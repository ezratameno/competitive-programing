package main

func main() {

}

func minNumberOperations(target []int) int {
	// if we only have 1 number the the number of inc is that
	// number.
	ans := target[0]
	for i := 1; i < len(target); i++ {
		// if the number is grater the the prev
		// than we add the diffrence
		if target[i] > target[i-1] {
			ans = ans + target[i] - target[i-1]
		}
	}

	return ans
}
