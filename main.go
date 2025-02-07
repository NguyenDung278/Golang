package main

import "fmt"

func main() {
	counter := 1
	for a := 1000; a <= 99929; a++ {
		for b := a; b <= 9999; b++ {
			s := fmt.Sprintf("%d", a*b)
			if s[0] == s[len(s)-1] {
				counter++
			}
		}
	}
	println(counter)
	fmt.Println(" 12121212wo212r122232222ld12")
}
