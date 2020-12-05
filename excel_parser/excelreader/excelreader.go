package excelreader

import (
	vars "../vars"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/360EntSecGroup-Skylar/excelize"
	_ "strconv"
	"strings"
)


func ConvertExcel2EntityList(path string) (error, []vars.CDREntity) {
	var f, err = excelize.OpenFile(path)
	cdrList := []vars.CDREntity{}
	if err != nil {
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
		header2Entity := vars.GetHeader2Entity()
		for _, colCell := range row {
			header := vars.Address2Header[cellAddress]
			field := header2Entity[header]
			vars.SetField(&cdrEntity, field, colCell)
			cellAddress = nextCell(cellAddress)
		}
		cdrList = append(cdrList, cdrEntity)
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
	return vars.ABC[position + 1: position + 2]
}

