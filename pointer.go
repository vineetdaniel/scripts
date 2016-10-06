package main
import "fmt"

func zero(x int) {
	x = 0
	fmt.Println(x)
}

func pZero(xPtr *int) {
	*xPtr = 2
}

func main() {
	x := 5
	y := 4
	zero(x)
	pZero(&y) 
	fmt.Println(x)
	fmt.Println(y)
}
