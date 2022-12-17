package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := 10
	b := "Hello"
	c := 10.44
	d := false
	e := 'w'
	f := `Gustavo
	      Lindo
	`
	res, err := http.Get("https://github.com/api")

	if err != nil {
		panic("ERRRROOO")
	}

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)

}
