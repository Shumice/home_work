package main

import (
	"fmt"
	"hw02_fix_app/printer"
	"hw02_fix_app/reader"
	"hw02_fix_app/types"
)

func main() {
	var path string = "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"

	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
