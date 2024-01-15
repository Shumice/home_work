package main

import "fmt"

func main() {
	var cell int
	fmt.Println("Введи колличество ячеек:")
	fmt.Scanf("%d\n", &cell)

	first := "#"
	second := " "

	for i := 1; i <= cell; i++ {
		if i%2 == 0 {
			for j := 1; j <= cell; j++ {
				if j%2 == 0 {
					outCell := first
					fmt.Print(outCell)
				}
				if j%2 != 0 {
					outCell := second
					fmt.Print(outCell)
				}
			}
		}
		if i%2 != 0 {
			for j := 1; j <= cell; j++ {
				if j%2 == 0 {
					outCell := second
					fmt.Print(outCell)
				}
				if j%2 != 0 {
					outCell := first
					fmt.Print(outCell)
				}
			}
		}
		fmt.Println()
	}

	// Place your code here.
}
