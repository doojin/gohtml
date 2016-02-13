package util

import (
	"sort"
)

// SortMapKeys returns sorted list o map keys
func SortedMapKeys(m map[string]interface{}) (keyList []string) {
	for key := range m {
		keyList = append(keyList, key)
	}
	sort.Strings(keyList)
	return
}
