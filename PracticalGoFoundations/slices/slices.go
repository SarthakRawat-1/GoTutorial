package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var s []int
	fmt.Println("len", len(s)) // len is "nil safe"

	if s == nil { // You can compare only a slice to nil
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4]

	// s3 -> [2 3 4], but it shares the *same underlying array* as s2
	// s3 has capacity = len(s2) - 1 = 6

	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:100])

	s3 = append(s3, 100)
	// Go checks if s3 has enough capacity to append without creating a new array.
	// Here:
	//     len(s3) = 3
	//     cap(s3) = 6 (because it’s slicing from index 1 of s2, and s2 had 7 elements)
	// Since cap(s3) is larger than len(s3), Go reuses the same underlying array, placing 100 right after the last element of s3 in that array.
	// That means the element at index 4 in the shared array (which corresponds to s2[4]) gets overwritten with 100.
	// So s2 changes too.

	fmt.Printf("s3 = %#v\n", s3)
	fmt.Printf("s2 = %#v\n", s2) // s2 also changed.

	var s4 []int

	for i := 0; i < 1_000_000; i++ {
		s4 = appendInt(s4, i)
	}

	fmt.Println("s4", len(s4), cap(s4))

	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"}))

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))

	fmt.Println(reflect.TypeOf(2))
}

func median(values []float64) float64 {
	// Copy in order not to change values
	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2

	if len(nums) % 2 == 1 {
		return nums[i]
	} else {
		return (nums[i - 1] + nums[i]) / 2
	}
}

func concat (s1, s2 []string) []string {
	s := make([]string, len(s1) + len(s2))

	copy(s, s1)
	copy(s[len(s1):], s2)

	return s
}

func appendInt (s []int, v int) []int {
	i := len(s)

	if len(s) < cap(s) { // Enough space in underlying array
		s = s[:len(s) + 1]
	} else { // Need to reallocate and copy
		fmt.Printf("Reallocate :- %d -> %d\n", len(s), 2 * len(s) + 1)
		s2 := make([]int, 2 * len(s) + 1)
		copy(s2, s)
		s = s2[:len(s) + 1]
	}

	s[i] = v

	return s
}