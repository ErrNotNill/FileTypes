package txt

import (
	"fmt"
	"os"
	"strings"
	"xml-txt/pkg"
)

var modifiedUserArgs = map[string]string{"1": "A", "2": "B", "3": "C", "4": "D", "5": "E", "6": "F", "7": "G", "8": "H", "9": "I", "10": "J",
	"11": "K", "12": "L", "13": "M", "14": "N", "15": "O", "16": "P", "17": "Q", "18": "R", "19": "S", "20": "T", "21": "Y", "22": "V",
	"23": "W", "24": "X", "25": "Y", "26": "Z",
}

var Args []string
var UserArgsCel string

func GetArgsFromTxt() {
	read, _ := os.ReadFile(pkg.TxtFile)

	split := strings.Split(string(read), ";")
	Args = append(split)

	fmt.Println(Args)

	// todo opts for user Args (example Args[0])
	setCell := Args[0:2]
	fmt.Println("SetCell is: ", setCell)
}

func ExcelRow() string {
	return ""
}
func ExcelColumn() string {
	return ""
}
func NextColumn() string {
	return ""
}

//opts

var Row string
var Column string

func swap() string {
	return Column + Row
}
