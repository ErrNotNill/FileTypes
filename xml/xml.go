package xml

import (
	"fmt"
	"github.com/beevik/etree"
	"xml-txt/pkg"
)

var Article = make([]string, 0) //here all Articles

var UserArgsArray []string

func ParseXml(scan string) {

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(pkg.XmlFile); err != nil {
		panic(err)
	}
	root := doc.SelectElement("data")
	fmt.Println("ROOT element:", root.Tag)

	//todo NEXT WORK HERE

	for _, book := range root.SelectElements("product") {

		if userargs := book.SelectElement(scan); userargs != nil { //todo User Args
			Article = append(Article, userargs.Text())
			fmt.Println(userargs.Text())
		}

	}
}
