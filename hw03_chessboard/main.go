package main

import "fmt"

func main() {
	var cell int
	fmt.Println("Введи колличество ячеек")
	fmt.Scanf("%d\n", &cell)

	for i := 1; i <= cell; i++ {
		if i%2 == 0 {
			for j := 1; j <= cell; j++ {
				fmt.Printf(" " + "#")
			}
		}
		if i%2 != 0 {
			for j := 1; j <= cell; j++ {
				fmt.Printf("#" + " ")
			}
		}
		fmt.Printf("\n")
	}

	// Place your code here.
}
