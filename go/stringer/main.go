package main

//https://last9.io/blog/golang-stringer-tool/
import "fmt"

type StatusCode int

const (
	statusOK                StatusCode = 200
	statusInternalServerErr StatusCode = 500
)

func main() {
	fmt.Println(statusOK)
	fmt.Println(statusInternalServerErr)

	// Output:
	// 200
	// 500
}
