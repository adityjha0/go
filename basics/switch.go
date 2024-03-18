package main

import (
	"fmt"
	"time"
)

func main(){
	i:=1
	fmt.Print(i, " is ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("either greater than three or less than one")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

	WhoAmI := func (i interface{})  {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("type %T\n", t)
		}
	}

	WhoAmI(1)
	WhoAmI(true)
	WhoAmI("shfsilfs")

}
