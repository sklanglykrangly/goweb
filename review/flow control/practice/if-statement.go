package main
import "fmt"


func main()  {
	x := 0

	for x = 0; x <=25; x++ {
		if x % 2 == 0 {
			fmt.Println(x, "is even")
		} else {
			fmt.Println(x, "is odd")
		}
	}
}



// within a loop, use a series of if statements demonstrating if / else-if / else-if / end 