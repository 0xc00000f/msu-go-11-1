package main

import (
	"sort"
	"strconv"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(slice []int) string {
	answer := ""
	for _, value := range slice {
		answer += strconv.Itoa(value)
	}
	return answer
}

func MergeSlices(firstSlice []float32, secondSlice []int32) []int {
	answer := make([]int, 0)
	for _, value := range firstSlice {
		answer = append(answer, int(value))
	}
	for _, value := range secondSlice {
		answer = append(answer, int(value))
	}

	return answer
}

func GetMapValuesSortedByKey(m map[int]string) []string {

	result := make([]string, 0)

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range keys {
		result = append(result, m[key])
	}

	return result
}

func main() {

}
