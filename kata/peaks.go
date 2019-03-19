package kata

import (
	"fmt"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// Not optimal
// PickPeaks: https://www.codewars.com/kata/5279f6fe5ab7f447890006a7/solutions/go
func PickPeaks(array []int) PosPeaks {
	peaks := &PosPeaks{Pos: []int{}, Peaks: []int{}}

	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i+1] && array[i] > array[i-1] {
			peaks.Pos = append(peaks.Pos, i)
			peaks.Peaks = append(peaks.Peaks, array[i])
		} else if array[i] == array[i+1] && array[i] > array[i-1] {

			fmt.Println(array[i])
			y := i
			for ; array[i] == array[y] && y < len(array)-1; y++ {
				if array[y] > array[y+1] {
					peaks.Pos = append(peaks.Pos, i)
					peaks.Peaks = append(peaks.Peaks, array[i])
				}
			}
			i = y - 1
		}
	}

	return *peaks
}
