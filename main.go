package main

import "fmt"

func main() {
	counter1 := 1
	for a := 1000; a <= 99929; a++ {
		for b := a; b <= 9999; b++ {
			s := fmt.Sprintf("%d", a*b)
			if s[0] == s[len(s)-1] {
				counter1++
			}
		}
	}
	println(counter1)
	fmt.Println(" 12121212wo212r122232222ld12")
}
