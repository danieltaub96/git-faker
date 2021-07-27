package array

import "math/rand"

func RandomValueFromArray(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	return arr[rand.Intn(len(arr))]
}
