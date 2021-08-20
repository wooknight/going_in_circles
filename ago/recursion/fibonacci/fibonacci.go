package main

import (
	"fmt"
	"time"

)

func fibonacci (num int) int{
	if num <= 1 {
		return 1
	}
	return fibonacci(num-2)+fibonacci(num-1)
}

func fibCaller(num int) {
	defer track (runningtime("fibonacci"))
	fmt.Println(fibonacci(num))

}

func main (){
	fibCaller(20)
	dynMapperCaller(40)
}

func dynMapperCaller(num int){
	defer track (runningtime("dynMapperCaller"))
	mapper:=make(map[int] int)
	mapper[0]=1
	mapper[1]=1
	fmt.Println(dynfibonacci(num,mapper))
}

func runningtime(s string) (string , time.Time) {
	fmt.Println("Start :",s)
	return s , time.Now()
}

func track(s string , startTime time.Time) {
	endTime:=time.Now()
	fmt.Println("End :",s,"took ",endTime.Sub(startTime))
}

func dynfibonacci (num int, fibMapper map[int] int) int{
	if num < 0 {
		return 1
	}
	if val, ok:= fibMapper[num]; !ok {
		fibMapper[num] = fibonacci(num-2)+fibonacci(num-1)
	}else{
		return val
	}
	return fibMapper[num]
}
