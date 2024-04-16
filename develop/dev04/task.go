package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"пятак", "лИсток", "пЯтка", "acbd", "тяпка", "слиток", "столик", "abcd"}
	fmt.Println(searchAnagram(words))
}

func searchAnagram(masWords []string) map[string][]string {
	anagrams := make(map[string][]string, 10)

	for _, s := range masWords {
		flag := false
		for key, _ := range anagrams {
			v, l := countLettersUnicode(key)
			v1, l1 := countLettersUnicode(s)
			if l == l1 {
				if equalMap(v, v1) {
					anagrams[key] = append(anagrams[key], s)
					flag = true
					break
				}
			}
		}
		if !flag {
			anagrams[s] = []string{}
		}
	}

	for k, v := range anagrams {
		if len(v) <= 1 {
			delete(anagrams, k)
		}
	}
	return anagrams
}

func countLettersUnicode(str string) (map[string]int, int) {
	countUniq := make(map[string]int, 5)
	chars := []rune(strings.ToLower(str))
	length := len(chars)
	for i := 0; i < length; i++ {
		countUniq[string(chars[i])] += 1
	}
	return countUniq, length
}

func equalMap(mapF map[string]int, mapS map[string]int) bool {
	for keyF, valF := range mapF {
		valS, ok := mapS[keyF]
		if ok {
			if valF != valS {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
