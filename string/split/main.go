package main

import (
	"fmt"
	"strings"
)

// You are given a string S consisting of letters 'a' and 'b'. You want to split it into three separate non-empty parts.
// The lengths of the parts can differ from one another.
// In how many ways can you split S into three parts, such that each part contains the same number of letters 'a'?

// Write a function:

// int solution(string &S);
// that, given a string S of length N, returns the number of possible ways of splitting S as described above.

// Examples:

// 1. Given S = "babaa", the function should return 2. The possible splits are: "ba | ba | a" and "bab | a | a".
// 2. Given S = "ababa", the function should return 4. The possible splits are: "a | ba | ba", "a | bab | a", "ab | a | ba" and "ab | ab | a".
// 3. Given S = "aba", the function should return 0. It is impossible to split S as required.
// 4. Given S = "bbbbb", the function should return 6. The possible splits are: "b | b | bbb", "b | bb | bb", "b | bbb | b", "bb | b | bb", "bb | bb | b" and "bbb | b | b". Each part contains the same number of letters 'a', i.e. 0.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..40,000];
// string S contains only letters 'a' and 'b'.
const (
	A          string = "a"
	SPLITCOUNT int    = 3
)

func main() {
	str := []string{"babaa", "ababa", "aba", "bbbbb"}
	for _, v := range str {
		fmt.Println(v, " :", solution(v))
	}

}
func solution(str string) int {
	// checking 'a' occurance
	countOfa := strings.Count(str, A)
	if countOfa == 0 {
		// if 0 that mean it eqaully distributed in all i.e 0
		return ((len(str) - 1) * (len(str) - 2)) / 2
	}
	// if a occurance is not divisble by SPLITCOUNT that mean a is not equally distributed
	if countOfa%SPLITCOUNT != 0 {
		return 0
	}

	acceptedCount := countOfa / SPLITCOUNT
	count, first, second := 0, 0, 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == A {
			count += 1
		}
		if count == acceptedCount {
			first += 1
		} else if count == 2*acceptedCount {
			second += 1
		}
		// fmt.Println("count: ", count, " first: ", first, " second: ", second)
	}
	return first * second
}
