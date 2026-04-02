package main

import "strings"

func fixPunctuation(s string) string {
	words := strings.Fields(s)
	for i := len(words) - 1; i >= 0; i-- {

		if strings.Contains(".,!?:;", words[i]) {
			if i > 0 {
				words[i-1] = words[i-1] + words[i]
			}
			words = append(words[:i], words[i+1:]...)
			continue
		}
	}
	for i := 0; i < len(words); i++ {
		if i > 0 {
			if words[i] == "..." {
				words[i-1] = words[i-1] + words[i]
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}
			if strings.HasPrefix(words[i], ",") {
				words[i-1] = words[i-1] + ","
				words[i] = words[i][1:]
			}
		}
	}
	return strings.Join(words, " ")
}
