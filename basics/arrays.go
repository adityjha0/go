package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("empty : ", a)
	
	a[4] = 100
	fmt.Println("set index 4 : ", a)
	fmt.Println("get index 4 : ", a[4])

	fmt.Println("length : ", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("declare : ", b)

	var array2D [2][3]int
	for i:=0; i<2; i++{
		for j:=0; j<3; j++{
			array2D[i][j] = i+j
		}
	}

	fmt.Println("2d array : ", array2D)
}
