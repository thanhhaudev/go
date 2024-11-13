package main

import "fmt"

func main() {
	fmt.Println("romanToInt:", romanToInt("MCMXCIV")) // 3
}

func romanToInt(s string) int {
	symbolMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var (
		res  int
		prev int
	)

	for i := len(s) - 1; i >= 0; i-- { // loop from right to left
		cur := symbolMap[s[i]]
		if cur >= prev {
			res += cur
		} else {
			res -= cur
		}

		prev = cur

		//fmt.Println(i, ". current", string(s[i]), "cur:", cur, "prev: ", prev, "res: ", res)
	}

	return res
}

//func romanToInt(s string) int {
//	symbolMap := map[byte]int{
//		'I': 1,
//		'V': 5,
//		'X': 10,
//		'L': 50,
//		'C': 100,
//		'D': 500,
//		'M': 1000,
//	}
//
//	var res int
//	for i := 0; i < len(s); i++ {
//		val := s[i]
//		if i < len(s)-1 && symbolMap[val] < symbolMap[s[i+1]] {
//			res += symbolMap[s[i+1]] - symbolMap[val]
//			i++
//		} else {
//			res += symbolMap[val]
//		}
//	}
//
//	return res
//}
