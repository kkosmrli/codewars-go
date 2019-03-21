package main

import (
	"fmt"
	"github.com/kkosmrli/codewars-go/kata"
)

func main() {
	// Rainfall kata
	//fmt.Println(kata.Mean("London", kata.DaTa))
	//fmt.Println(kata.Variance("London", kata.DaTa))

	// Non repeating char
	//fmt.Println(kata.FirstNonRepeating("stress"))

	// Maximum subarray
	//fmt.Println(kata.MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	// peaks.go
	//fmt.Println(kata.PickPeaks([]int{2, 1, 3, 2, 2, 2, 2, 1}))

	// var test = [][]int{
	// 	{7, 3, 5, 8, 1, 9, 6, 4, 2},
	// 	{1, 9, 4, 2, 6, 5, 8, 7, 3},
	// 	{3, 1, 7, 5, 8, 4, 2, 6, 9},
	// 	{6, 5, 9, 1, 7, 2, 4, 3, 8},
	// 	{4, 8, 2, 9, 3, 6, 7, 1, 5},
	// 	{9, 4, 8, 7, 5, 1, 3, 2, 6},
	// 	{5, 6, 1, 4, 2, 3, 9, 8, 7},
	// 	{2, 7, 3, 6, 9, 8, 1, 5, 4},
	// 	{1, 2, 6, 3, 4, 7, 5, 9, 8},
	// }
	// var test = [][]int{
	// 	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	// 	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	// 	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	// 	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	// 	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	// 	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	// 	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	// 	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	// 	{3, 4, 5, 2, 8, 6, 1, 7, 9},
	// }
	// fmt.Println(kata.ValidateSolution(test))
	s1 := "Program title: Primes\nAuthor: Kern\nCorporation: Gold\nPhone: +1-503-555-0071\nDate: Tues April 9, 2005\nVersion: 6.7\nLevel: Alpha"
	s11 := "Program title: Hammer\nAuthor: Tolkien\nCorporation: IB\nPhone: +1-503-555-0090\nDate: Tues March 29, 2017\nVersion: 2.0\nLevel: Release"
	s12 := "Program title: Primes\nAuthor: Kern\nCorporation: Gold\nPhone: +1-503-555-009\nDate: Tues April 9, 2005\nVersion: 6.7\nLevel: Alpha"

	fmt.Println(kata.Change(s1, "Ladder", "1.1"))
	fmt.Println(kata.Change(s11, "Balance", "1.5.6"))
	fmt.Println(kata.Change(s12, "blah", "blub"))

}
