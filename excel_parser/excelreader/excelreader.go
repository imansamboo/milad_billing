package excelreader

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	_ "strconv"
	vars "../vars"
	"strings"
)


func ConvertExcel2EntityList(path string) (error, []vars.CDREntity) {
	var f, err = excelize.OpenFile(path)
	cdrList := []vars.CDREntity{}
	if err != nil {
		fmt.Println("not valid file")
		fmt.Println(err)
		return err, nil
	}
	cnt := 0
	// Get value from cell by given worksheet name and axis.
	cellAddress := "A"
	//Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic("some thing wrong in reading all rows")
	}
	for _, row := range rows {
		cnt += 1
		if cnt < 4 {
			continue
		}
		cdrEntity := vars.CDREntity{}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
			field := vars.Address2Header[cellAddress]
			settField(&cdrEntity, field, colCell)
			cellAddress = nextCell(cellAddress)
		}
		cdrList = append(cdrList, cdrEntity)
		fmt.Println()
	}
	return nil, cdrList

	//cell, err := f.GetCellValue("Sheet1", "A3")
	//if err != nil {
	//	fmt.Println(err)
	//	return err, nil
	//}
	//fmt.Println(cell)
	////Get all the rows in the Sheet1.
	//rows, err := f.GetRows("Sheet1")
	//if err != nil {
	//	panic("some thing wrong")
	//}
	//cnt := 0
	//for _, row := range rows {
	//	if cnt < 2 {
	//		continue
	//	}
	//	cnt += 1
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}
}

func nextCell(cellAddress string) string{
	if cellAddress == "W"{
		return "A"
	}
	position := strings.Index(vars.ABC, cellAddress)
	return vars.ABC[position: position + 1]
}

func settField(v *vars.CDREntity, field string, value string) {
	r := reflect.ValueOf(v).Elem()
	r.FieldByName(field).SetString(value)
}
