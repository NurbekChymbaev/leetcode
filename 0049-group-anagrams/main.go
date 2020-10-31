package _049_group_anagrams

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	var resultString [][]string
	var result = make(map[int][]string)
	var storage = make(map[string]int)
	var storageIndex = 0

	for i := range strs {

		key := sortString(strs[i])

		if _, ok := storage[key]; ok {
			result[storage[key]] = append(result[storage[key]], strs[i])
		} else {
			result[storageIndex] = append(result[storageIndex], strs[i])
			storage[key] = storageIndex
			storageIndex++
		}
	}

	for _, k := range result {
		resultString = append(resultString, k)
	}

	return resultString
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
