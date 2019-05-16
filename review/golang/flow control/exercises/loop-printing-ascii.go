package main
import "fmt"


func main()  {
	x := 33

	for x = 33; x <=122; x++ {
		fmt.Printf("%v\t%#U\t%+q\n\n", x,x,x)
	}
}


// create a loop that prints all upper and lower case ascii letter values - from 33 to 122