package function

import "fmt"

func FizzBuzz(slice []int) {
	for i := 0; i < len(slice); i++ {
		if slice[i]%15 == 0 {
			fmt.Println("fizzBuzz")
		} else if slice[i]%5 == 0 {
			fmt.Println("buzz")
		} else if slice[i]%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(slice[i])
		}
	}
}
