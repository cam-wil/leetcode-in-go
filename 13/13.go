package main

import (
    "fmt"
)

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
    
	sum := romanMap[s[0]]
    
	for i := 1; i < len(s); i++ {
        roman := s[i]
		prev := romanMap[s[i-1]]
		if prev < romanMap[roman] {
			sum -= (prev * 2)
		}
		sum += romanMap[roman]
	}
	return sum
}

func main() {
    fmt.Println("X: \t", romanToInt("X"))
    fmt.Println("XI: \t", romanToInt("XI"))
    fmt.Println("IX: \t", romanToInt("IX"))
}
