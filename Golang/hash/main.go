// Isomorphic Strings - Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
//  All occurrences of a character must be replaced with another character while preserving
// the order of characters. No two characters may map to the same character,
//  but a character may map to itself. s and t consist of any valid ascii characte

package main

import "fmt"

// T(o)=n^2
// S(o)=1

func isIsomorpichByHashMap(s string, t string) bool {
	sHash := map[byte]byte{}
	tHash := map[byte]byte{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		if mappedS, exist := sHash[charS]; exist {
			if mappedS != charT {
				return false
			}
		} else {
			sHash[charS] = charT
		}

		if mappedT, exist := tHash[charT]; exist {
			if mappedT != charS {
				return false
			}
		} else {
			tHash[charT] = charS
		}
	}
	return true
}
func isIsomorphicByBeruteForcer(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && t[i] != t[j] {
				return false
			}
			if t[i] == t[j] && s[i] != s[j] {
				return false
			}
		}
	}

	return true
}
func main() {
	s := "Herto"
	t := "Wossd"
	brt := isIsomorphicByBeruteForcer(s, t)

	nbrt:=isIsomorpichByHashMap(s, t)
	fmt.Println(brt)
	fmt.Println(nbrt)

}
