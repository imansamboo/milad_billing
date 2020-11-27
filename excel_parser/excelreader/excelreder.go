package excelreader

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/360EntSecGroup-Skylar/excelize"
)


func ConvertExcel2Map(path string) {
	var f, err = excelize.OpenFile(path)
	if err != nil {
		fmt.Println("not valid file")
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "A3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	//Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic("some thing wrong")
	}
	cnt := 0
	for _, row := range rows {
		if cnt < 2 {
			continue
		}
		cnt += 1
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
