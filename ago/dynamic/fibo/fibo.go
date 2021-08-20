package main

import (
	"math"
)

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
	// fmt.Println(climber(n))
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
	memo := make([][]int, cnt+1)
	for i := 0; i < cnt+1; i++ {
		memo[i] = make([]int, choose+1)
	}

	var countSubsetsMemoHelper func(cnt, choose int) int
	memo[1][0] = 1
	memo[1][1] = 1
	countSubsetsMemoHelper = func(cnt, choose int) int {
		if cnt == choose || cnt == 0 || choose == 0 {
			memo[cnt][choose] = 1
		}
		if memo[cnt][choose] == 0 {
			// memo[cnt-1][choose] = countSubsetsMemoHelper(cnt-1, choose)
			// memo[cnt-1][choose-1] = countSubsetsMemoHelper(cnt-1, choose-1)
			// memo[cnt][choose] = memo[cnt-1][choose] + memo[cnt-1][choose-1]
			memo[cnt][choose] = countSubsetsMemoHelper(cnt-1, choose) + countSubsetsMemoHelper(cnt-1, choose-1)
		}
		return memo[cnt][choose]
	}
	return countSubsetsMemoHelper(cnt, choose)
}

func countSubsetsMemoMap(cnt, choose int) int {
	memo := make(map[[2]int]int)
	var countSubsetsMemoMapHelper func(cnt, choose int) int
	memo[[2]int{0, 0}] = 1
	memo[[2]int{1, 1}] = 1
	countSubsetsMemoMapHelper = func(cnt, choose int) int {
		if cnt == choose || cnt == 0 || choose == 0 {
			memo[[2]int{cnt, choose}] = 1
		}
		if _, ok := memo[[2]int{cnt, choose}]; !ok {
			// memo[cnt-1][choose] = countSubsetsMemoHelper(cnt-1, choose)
			// memo[cnt-1][choose-1] = countSubsetsMemoHelper(cnt-1, choose-1)
			// memo[cnt][choose] = memo[cnt-1][choose] + memo[cnt-1][choose-1]
			memo[[2]int{cnt, choose}] = countSubsetsMemoMapHelper(cnt-1, choose) + countSubsetsMemoMapHelper(cnt-1, choose-1)
		}
		return memo[[2]int{cnt, choose}]
	}
	return countSubsetsMemoMapHelper(cnt, choose)
}

func countSubsetsPregen2DMemo(cnt, choose int) int {
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

func countPaths(row, col int) int {
	table := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		table[i] = make([]int, col+1)
	}
	for i := 0; i <= col; i++ {
		table[0][i] = 1
	}
	for i := 0; i <= row; i++ {
		table[i][0] = 1
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			table[i][j] = table[i-1][j] + table[i][j-1]
		}
	}
	// fmt.Printf("%+v\n\n", table)
	return table[row-1][col-1]
}

//optimize for cost

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCost(len int, costarray []int) int {
	table := make([]int, len+2)
	table[0] = 0
	table[1] = costarray[0]
	costarray = append(costarray, 0)
	for i := 2; i <= len+1; i++ {
		table[i] = costarray[i-1] + min(table[i-1], table[i-2])
	}
	return table[len+1]
}

func coinMinChange(amount int, coins []int) int {
	table := make([]int, amount+1)
	table[0] = 0
	for i := 1; i < amount+1; i++ {
		table[i] = int(math.MaxInt32)
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c < 0 {
				continue
			}
			if table[i-c] < table[i] {
				table[i] = table[i-c]
				table[i] = table[i] + 1
			}
		}
	}

	return table[amount]
}
