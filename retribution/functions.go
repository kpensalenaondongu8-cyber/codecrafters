package main

import (
	"strconv"
	"strings"
)

func processCommands(s string) string {
	words := strings.Fields(s)

	for i := len(words) - 1; i >= 0; i-- {
		if strings.HasPrefix(words[i], "(up,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)

			for j := 1; j <= n; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			continue
		}

		if strings.HasPrefix(words[i], "(low,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)

			for j := 1; j <= n; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToLower(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			continue
		}
		if strings.HasPrefix(words[i], "(cap,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)

			for j := 1; j <= n; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(string(words[i-j][0])) + strings.ToLower(words[i-j][1:])
				}
			}
			words = append(words[:i], words[i+2:]...)
			continue
		}

		if i > 0 {
			if words[i] == "(hex)" || words[i] == "(bin)" || words[i] == "(cap)" || words[i] == "(up)" || words[i] == "(low)" {
				if words[i] == "(hex)" {
					num, _ := strconv.ParseInt(words[i-1], 16, 64)
					words[i-1] = strconv.FormatInt(num, 10)
				}
				if words[i] == "(bin)" {
					num, _ := strconv.ParseInt(words[i-1], 2, 64)
					words[i-1] = strconv.FormatInt(num, 10)
				}
				if words[i] == "(up)" {
					words[i-1] = strings.ToUpper(words[i-1])
				}
				if words[i] == "(low)" {
					words[i-1] = strings.ToLower(words[i-1])
				}
				if words[i] == "(cap)" {
					word := words[i-1]
					if len(word) > 0 {
						words[i-1] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
					}
				}
				words = append(words[:i], words[i+1:]...)
				continue
			}
		}

	}
	return strings.Join(words, " ")
}
