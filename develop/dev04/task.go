package main

import (
	"fmt"
	"sort"
	"strings"
)

// структура с мапой анаграмм и вспомогательной
// мапой для поиска подходящего ключа в основной
type anagrams struct {
	words map[string][]string
	aux   map[string]string
}

func main() {
	a := mapAnagrams([]string{"пятак", "сорт", "столик", "ток",
		"листок", "рост", "ватка", "тяпка", "торс", "кот", "пятка", "трос", "слиток"})
	fmt.Println(a.findAnagrams("Рост"))
}

// наполнение мапы анаграммами
func mapAnagrams(words []string) *anagrams {
	aMap := &anagrams{words: make(map[string][]string), aux: make(map[string]string)}
	for _, wordStr := range words {
		wordStr = strings.ToLower(wordStr)
		wordSorted := sortRunes([]rune(wordStr))
		mKey, ok := aMap.aux[string(wordSorted)]
		if !ok {
			mKey = wordStr
			aMap.aux[string(wordSorted)] = wordStr
		}
		aMap.insertWord(mKey, wordStr)
	}
	aMap.deleteSingles()
	return aMap
}

func (aMap *anagrams) insertWord(key string, word string) {
	words, _ := aMap.words[key]
	i, found := findPlace(word, words)
	if found {
		return
	}
	if len(words) == 0 || len(words) == i {
		words = append(words, word)
	} else {
		tmp := make([]string, i)
		copy(tmp, words[:i+1])
		tmp = append(tmp, word)
		tmp = append(tmp, words[i:]...)
		words = tmp
	}
	aMap.words[key] = words
}

func (aMap *anagrams) deleteSingles() {
	for key, words := range aMap.words {
		if len(words) == 1 {
			delete(aMap.words, key)
		}
	}
}

// поиск индекса для вставки нового элемента в отсортированном массиве
func findPlace(word string, words []string) (int, bool) {
	low := 0
	high := len(words) - 1

	for low <= high {
		median := low + (high-low)/2

		if words[median] < word {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(words) || words[low] != word {
		return low, false
	}

	return 0, true
}

func sortRunes(r []rune) []rune {
	sort.Slice(r[:], func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}

func (aMap *anagrams) findAnagrams(word string) []string {
	word = strings.ToLower(word)
	wordSorted := string(sortRunes([]rune(word)))
	key, ok := aMap.aux[wordSorted]
	if ok {
		return aMap.words[key]
	}
	return nil
}
