package main
import "fmt"


func main()  {
	i := 0
	j := 0
	for i = 0; i <= 10; i++ {
			// fmt.Println(i)
			for j = 0; j < 2; j++ {
				fmt.Println(i)
			}
	}
}