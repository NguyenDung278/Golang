package main

import "fmt"

func main() {
	counter := 1
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			s := fmt.Sprintf("%d", a*b)
			if s[0] == s[len(s)-1] {
				counter++
			}
		}
	}
	println(counter)
	fmt.Println(" wo212r122232222ld12")
	fmt.Println(counter)
}
