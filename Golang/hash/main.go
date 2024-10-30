// Isomorphic Strings - Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
//  All occurrences of a character must be replaced with another character while preserving
// the order of characters. No two characters may map to the same character,
//  but a character may map to itself. s and t consist of any valid ascii character.

package main

import "fmt"

func isIsomorpic(s string, t string) bool {

	sHash := map[byte]byte{}
	tHash := map[byte]byte{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		if mappedCharS, exist := sHash[charS]; exist {
			if mappedCharS != charT {
				return false
			}
		} else {
			sHash[charS] = charT
		}
		if mappedCharT, exist := tHash[charT]; exist {
			if mappedCharT != charS {
				return false
			}
		} else {
			tHash[charT] = charS
		}

	}
	return true
}

func main() {
	s := "Hello"
	t := "Wolld"

	iso := isIsomorpic(s, t)
	fmt.Println(iso)
}
