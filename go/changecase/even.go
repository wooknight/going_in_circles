package main

import "fmt"

func countEven(num int) int {
	cnt := 0
	calcs := make(map[int]bool)
	for i := 1; i < num; i++ {
		if digitsum(i, calcs) == true {
			fmt.Println(i)
			calcs[i] = true
			cnt++
		}
	}
	return cnt
}

func digitsum(num int, calcs map[int]bool) bool {
	// if calcs[num] == true{
	//     return true
	// }
	if num/10 == 0 {
		calcs[num] = (num % 2) == 0
		return calcs[num]
	}
	var sum int
	for num/10 > 0 {
		sum += num % 10
		num = num / 10
	}
	sum += num
	return sum%2 == 0
}

func main() {
	countEven(70)
}
