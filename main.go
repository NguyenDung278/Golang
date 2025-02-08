package main

import "fmt"

func main() {
	counter3 := 10
	for a := 1000; a <= 99929; a++ {
		for b := a; b <= 9999; b++ {
			s := fmt.Sprintf("%d", a*b)
			if s[0] == s[len(s)-1] {
				counter3++
			}
		}
	}
	fmt.Println(" 212w121o21212212r12122232222ld12")

	println(counter3)
	fmt.Println(" 212w121o21212212r12122232222ld12")
	println(counter3)
	fmt.Println(10)

	fmt.Println(" 1212121121212wo212r122232222ld12")
	fmt.Println(" 212w121o21212212r12122232222ld12")
}
