package array

import "math/rand"

func RandomValueFromArray(arr []string) string {
	return arr[rand.Intn(len(arr))]
}
