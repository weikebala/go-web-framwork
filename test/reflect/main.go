package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)
type T struct {
	Name string
	Age  int
	City string
}
func main() {
	x := T{"Michał", 99, "London"}
	y := T{"Adam", 88, "London"}
	if diff := cmp.Diff(x, y); diff != "" {
		fmt.Println(diff)
	}
}