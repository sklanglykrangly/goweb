package main
import "fmt"


func main()  {
	x := 10;
	// for x > 0 {
	// 	fmt.Println("...",x,"!")
	// 	x--
	// }
	// fmt.Println("lift off!")

	for {
		if x < 1 {
			break
		}
		fmt.Println("...",x,"!")
		x--
	}
	fmt.Println("lift off!")

}