package main
import "fmt"

// make a type 'cat'
type cat struct {
	cname string
}

// make a type 'catDetective'
type catDetective struct {
	cat
	licenseToKill bool

}

// make a function that uses type person to say something, function name 'speak'
func (p cat) speak()  {
	fmt.Println(p.cname, "says, 'Hello, I'm Acey Ace.'")
}

func (p catDetective) speak()  {
	fmt.Println(p.cname, "says, 'And I am JesseVentura, cat Detective.'")	
}

func main()  {
	// make a slice
	xi := []int{2,34,5,6,53,2,2}
	fmt.Println(xi)

	// map 
	mi := map[string]int {
		"Scott": 53,
		"Ace": 4,
	}
	fmt.Println(mi)

	// make a version of type person
	c1 := cat {
		"Ace the Cat",
	}
	fmt.Println(c1)
	cd := catDetective {
		cat {
			"JesseVentura",
		},
		true,
	}
	fmt.Println(cd)
	// call speak function with variable
	c1.speak()
	cd.speak()
}
