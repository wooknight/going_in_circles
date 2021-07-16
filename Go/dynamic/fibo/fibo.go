package main

import "fmt"

func memofibo(n int) int {
	memo := make(map[int]int)
	var fibo func(n int) int
	memo[0] = 0
	memo[1] = 1

	fibo = func(n int) int {
		if n == 0 || n == 1 {
			return n
		} else {
			if _, ok := memo[n]; !ok {
				memo[n] = fibo(n-1) + fibo(n-2)
			} else {
			}
			return memo[n]
		}
	}
	return fibo(n)
}

func fibo(n int) int {
	var fibo func(n int) int

	fibo = func(n int) int {
		if n == 0 || n == 1 {
			return n
		} else {
			return fibo(n-1) + fibo(n-2)
		}
	}
	return fibo(n)
}

func climbStairs(n int) int {
	var climber func(n int) int
	climber = func(n int) int {
		if n == 0 || n == 1 || n == 2 {
			return n
		}
		return climber(n-2) + climber(n-2)
	}
	fmt.Println(climber(n))
	return climber(n)
}

func makeChange(target int, coins []int) int {
	return 0
}

func countSubsets(cnt, choose int) int {
	if cnt == 0 || choose == 0 || cnt == choose {
		return 1
	}
	return countSubsets(cnt-1, choose) + countSubsets(cnt-1, choose-1)
}

func countSubsetsMemo(cnt, choose int) int {
	memo := make(map[[2]int]int)
	var countSubsetsMemoHelper func(cnt, choose int) int
	memo[[2]int{1, 0}] = 1
	memo[[2]int{1, 1}] = 1

	countSubsetsMemoHelper = func(cnt, choose int) int {
		if cnt == choose {
			memo[[2]int{cnt, choose}] = 1
			return memo[[2]int{cnt, choose}]
		}
		if _, ok := memo[[2]int{cnt, choose}]; !ok {
			memo[[2]int{cnt, choose}] = countSubsets(cnt-1, choose) + countSubsets(cnt-1, choose-1)
		}
		return memo[[2]int{cnt, choose}]
	}
	return countSubsetsMemoHelper(cnt, choose)
}

func countSubsets2DMemo(cnt, choose int) int {
	if choose == 0 || cnt == choose {
		return 1
	}
	memo2D := make([][]int, cnt+1)
	for i := 0; i < cnt+1; i++ {
		memo2D[i] = make([]int, choose+1)
	}
	for i := 0; i < cnt+1; i++ {
		//left most column
		memo2D[i][0] = 1
	}
	for i := 0; i <= choose; i++ {
		//diagonal cnt == choose
		memo2D[i][i] = 1
	}

	for row := 2; row < cnt; row++ {
		for col := 1; col <= row && col <= choose; col++ {
			if memo2D[row][col] != 1 {
				memo2D[row][col] = memo2D[row-1][col] + memo2D[row-1][col-1]
			}
		}
	}
	var helper func(cnt, choose int) int
	helper = func(cnt, choose int) int {
		if memo2D[cnt][choose] != 0 {
			return memo2D[cnt][choose]
		}
		memo2D[cnt-1][choose] = helper(cnt-1, choose)
		memo2D[cnt-1][choose-1] = helper(cnt-1, choose-1)
		return memo2D[cnt-1][choose] + memo2D[cnt-1][choose-1]
	}
	// fmt.Println(memo2D)
	return helper(cnt, choose)
}
