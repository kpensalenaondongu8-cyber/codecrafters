package main

import "strings"

func fixGrammar(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && isVowel(string(words[i+1][0])) {
			if words[i] == "a" && isLower(string(words[i+1][0])) {
				words[i] = "an"
			} else {
				words[i] = "An"
			}
		}
	}
	return strings.Join(words, " ")
}
