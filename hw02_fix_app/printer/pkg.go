package printer

import (
	"fmt"

	"github.com/Shumice/home_work/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	//var str string
	for i := 0; i < len(staff); i++ {
		//str := fmt.Sprintf("UserID: %d; Age: %d; Name: %s; DepartmentID: %d; ",
		//	staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(staff[i])
	}
	//fmt.Println(str)
}
