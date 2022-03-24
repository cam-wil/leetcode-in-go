package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var str string = ""
	var count = len(strs[0])

	for i := 1; i < len(strs); i++ { // loop over words
		var tempCount = 0 // store temp value for count
		for x := 0; x < len(strs[i]); x++ {
			if x >= len(strs[i-1]) || strs[i-1][x] != strs[i][x] {
				break
			}
			tempCount++

		}

		if tempCount < count { // if tempCount is smaller than count, count = tempCount
			count = tempCount
		}
	}

	for i := 0; i < count; i++ { // make string
		str += string(strs[0][i])
	}

	return str
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}
