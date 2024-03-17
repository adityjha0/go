package main

import "fmt"

func main() {

	fmt.Println("First")
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

	fmt.Println("\nSecond")
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

	fmt.Println("\nThird")
    for i := range 3 {
        fmt.Println("range", i)
    }

	fmt.Println("\nFourth")
    for {
        fmt.Println("loop")
        break
    }

	fmt.Println("\nFifth")
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}