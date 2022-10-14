package main

import (
	"fmt"
	"xml-txt/pkg"
)

func main() {
	readdir := new(pkg.Dir)
	file, _, _ := readdir.NewReadFilesFromDir(".")
	fmt.Println(file)

	rft := new(pkg.ReadFileTypes)
	fileContain := rft.FileContains("products (1605).xml", ".xml")
	fmt.Println(fileContain)
	/*if err := pkg.FindAllTypesInDir(); err != nil {
		fmt.Println(err.Error())
	}
	xml.ParseXml("article")  //opts what you need find
	xls.ParseExcel("033536") //todo method for user args*/

	//xls.GetColumnsExcel()
	//xls.ParseExcel("article")
	//txt.GetArgsFromTxt()

}
