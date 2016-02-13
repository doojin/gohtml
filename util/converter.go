package util

import (
	"strings"
	"strconv"
)

// IntSliceToString converts a slice of integers to string
func IntSliceToString(slice []int) string {
	stringSlice := []string{}
	for _, i := range slice {
		stringSlice = append(stringSlice, strconv.Itoa(i))
	}
	return strings.Join(stringSlice, " ")
}
