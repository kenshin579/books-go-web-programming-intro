package main

import "fmt"

func ExampleSlice_Append1() {
	ns1 := []int{1, 2, 3}
	ns1 = append(ns1, 4, 5)
	fmt.Println(ns1)

	//Output: [1 2 3 4 5]
}

func ExampleSlice_Append_Three_Dots() {
	ns1 := []int{1, 2, 3}
	ns2 := []int{6, 7, 8}
	ns1 = append(ns1, ns2...)
	fmt.Println(ns1)

	//Output: [1 2 3 6 7 8]
}

func ExampleSlice_Append3() {
	ns1 := []int{1, 2, 3}
	ns2 := []int{6, 7, 8}
	ns3 := []int{8, 9, 10, 11}
	ns1 = append(ns1, 4, 5)        // ns1: [1 2 3 4 5]
	ns1 = append(ns1, ns2...)      // ns1: [1 2 3 4 5 6 7 8]
	ns1 = append(ns1, ns3[1:3]...) // ns1: [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(ns1)

	//Output: [1 2 3 4 5 6 7 8 9 10]
}
