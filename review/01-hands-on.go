package main
import "fmt"

// create a type square
type square struct {

}

type circle struct {
	pi int
	radius int
}

func (c circle) area() {
	radi := c.radius
	radi := 8
	a *= (c.pi * c.radius)
	fmt.Println("Area of circle is ", )
}

func main()  {
	fmt.Println()
}



// HANDS ON 1
// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle
https://play.golang.org/p/1enChb7Kg5 