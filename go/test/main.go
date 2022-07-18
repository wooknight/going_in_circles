package main 

import(
	"math"
	"fmt"
)

func oddEvenJumps(arr []int) int {
    jmps := 0    
    for i,val := range arr {
		if i == len(arr)-1{
			jmp++
		}else{
			if CanEnd(arr,i+1, val) {
				jmps++
			}	
		}
    }
    
    return jmps 
}

func CanEnd(arr []int,startIdx, val int) bool {
    jmps := 1
    end:=len(arr)
    
    for startIdx<end{
        if jmps % 2 == 1 {
            startIdx = getMinOfGreater (arr , startIdx , arr[startIdx])
            if startIdx == -1{
                break
            }
        }else{
            startIdx = getMaxofSmaller (arr , startIdx , arr[startIdx])
            if startIdx == -1{
                break
            }            
        }
        jmps++
    }
        return startIdx == end
}

func getMinOfGreater (arr []int, startIdx , val int) int {
    min := math.MaxInt
    next:=-1
    for i:=startIdx;i < len(arr);i++ {
        if arr[i]>val && arr[i]<min{
            min = arr[i]
            next = i
        }
    }
    return next
}

func getMaxofSmaller (arr []int, startIdx , val int) int {
    max := 0
        next:=-1
    for i:=startIdx;i < len(arr);i++ {
        if arr[i]<val && arr[i]>max{
            max = arr[i]
            next = i
        }
    }
    return next

}

func main (){
	// fmt.Println(oddEvenJumps([]int{10,13,12,14,15}))
	fmt.Println(oddEvenJumps([]int{14,15}))
}