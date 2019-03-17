package kata

import "math"

func WallPaper(l, w, h float64) string {
	num := map[float64]string{
		1.0:  "one",
		2.0:  "two",
		3.0:  "three",
		4.0:  "four",
		5.0:  "five",
		6.0:  "six",
		7.0:  "seven",
		8.0:  "eight",
		9.0:  "nine",
		10.0: "ten",
		11.0: "eleven",
		12.0: "twelve",
		13.0: "thirteen",
		14.0: "fourteen",
		15.0: "fifteen",
		16.0: "sixteen",
		17.0: "seventeen",
		18.0: "eighteen",
		19.0: "nineteen",
		20.0: "twenty"}
	if l*w*h == 0 {
		return "zero"
	}
	return num[math.Floor(((l*h*2.0+w*h*2.0)/(10.0*0.52))*1.15)+1.0]
}
