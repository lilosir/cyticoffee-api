package utils

import (
	"strconv"
	"strings"
)

// Int64ArrayToString converts an int64 array to a string
func Int64ArrayToString(arr []int64) string {
	res := make([]string, len(arr))
	for i, n := range arr {
		res[i] = strconv.FormatInt(n, 10)
	}
	return strings.Join(res, ",")
}

// IntArrayToString converts an int64 array to a string
func IntArrayToString(arr []int) string {
	res := make([]string, len(arr))
	for i, n := range arr {
		res[i] = strconv.Itoa(n)
	}
	return strings.Join(res, ",")
}
