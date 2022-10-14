package xls

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
	"xml-txt/pkg"
	"xml-txt/xml"
)

type Table struct {
	Col       string
	Row       string
	ClientCol string
	ClientRow string
}

func (t *Table) SplitTable() {
	t.ClientCol = t.Row
	t.ClientRow = t.Col
}

var FoundedArticle string

func GetColumnsExcel() {
	f, err := excelize.OpenFile(pkg.ExcelFile)
	if err != nil {
		fmt.Println(err)

	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("Остатки")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Printf("%T", colCell)
			//fmt.Print(colCell, "\t")

		}
		fmt.Println()
	}
}

func ParseExcel(scan string) bool {

	f, _ := excelize.OpenFile(pkg.ExcelFile)
	defer f.Close()

	rows, _ := f.GetRows("Остатки")

	//todo implement values in cells
	//setval := f.SetCellValue("Sheet2", "A2", "Hello world.")

	for _, v := range xml.Article { //todo find there Article

		if scan == v {
			fmt.Println(v)
			fmt.Printf("Yes i found this: %v", v)
			fmt.Println()
			FoundedArticle = v
			//todo find this article in Excel

			//opts return true
		}
	}
	for _, row := range rows {
		//row[] is long of array (rows), can be indexed to find column that you need, example row[0:len(row)+1]

		for _, colCell := range row {
			if strings.Contains(colCell, FoundedArticle) {
				fmt.Print("I found article in colCell from XLS: ", row)
			}

			//fmt.Print(colCell)
			cellArr = append(cellArr, colCell)

		}

	}
	return false
}

var cellArr []string
