package main
import "fmt"

func main()  {
	// set a variable value. use break and continue to check the value of the variable.
	// use a for statement to (de)increment the value. Use continue to check if the value still meets the
	// condition to loop. Use break when the value exceeds the condition of the loop

	// set variable 
	x := 10
	// establish for statement
	for {
		x--
		if x < 1 {
			break
		}
		if x % 2 == 0 {
			continue
		}
	}
}