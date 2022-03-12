package main
import (
	"fmt"
	"unicode"
	"strings"
)


func generateCase(str string) []string{
	if len(str) > 1 {
		if unicode.IsLetter(str[0]){
			str[0] = unicode.ToLower(str[0])
			strs1:=generateCase(str[1:])
			for idx,val:=range strs1{
				strs1[idx]= (string)str[0] + val
			}
			str[0] = unicode.ToUpper(str[0])
			strs2:=generateCase(str[1:])
			for idx,val:=range strs2{
				strs2[idx]= (string)str[0] + val
			}

			return append(strs1,strs2...)

		}else{
			strs1:=generateCase(str[1:])
			for idx,val:=range strs1{
				strs1[idx]= (string)str[0] + val
			}

		}
	}


	generateCase(str[1:])
}


func permutate(nums []int) [][]int{
	if len(nums)<=1{
		return [][]{nums}
	}
	for i,val:=range nums{
		//keep first elem constant
		nums[1],nums[i] = nums[i],nums[1]
		arrs:= permutate(nums[1:])
		for i, res:= range arrs{
			arrs[i]= append(arrs)
		}
	}
}