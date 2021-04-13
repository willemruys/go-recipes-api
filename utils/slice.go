package utils

import "github.com/lib/pq"

func IndexOfItemInSlice(slice pq.Int64Array, value int64) int {

	for p, v := range slice {
		if (v == value) {
			return p
		}
	}
	return -1
}

func RemoveItemFromSlice(s pq.Int64Array, i int) pq.Int64Array {
    s[i] = s[len(s)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return s[:len(s)-1]
}