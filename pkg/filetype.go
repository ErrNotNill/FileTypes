package pkg

import (
	"fmt"
	"github.com/h2non/filetype"
	"io/fs"
	"io/ioutil"
	"regexp"
	"strings"
)

type ReadFileTypes struct {
	Contains
}

type Contains interface {
	MimeFromFile(b []byte) string
	FileContains(fileName string, fileSuffix string) string
	FileSuffix(fileName string, fileSuffix string) string
	FileRegexp(contains string, fileName string) string
}

func MimeFromFile(b []byte) string {
	kind, _ := filetype.Match(b)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
	}
	return kind.Extension
}

//FileContains example contains == ".xls" but look in full string of fileName
func FileContains(fileName string, contains string) string {
	if strings.Contains(fileName, contains) {
		return fileName
	}
	return ""
}

//FileSuffix example fileSuffix == ".xml"
func FileSuffix(fileName string, fileSuffix string) string {
	if strings.HasSuffix(fileName, fileSuffix) {
		return fileName
	}
	return ""
}

type Mime interface {
}

//FileRegexp example contains == ".txt"
func FileRegexp(contains string, fileName string) string {
	r, _ := regexp.MatchString(contains, fileName)
	if r == true {
		return fileName
	}
	return ""
}

type Dir struct {
	directory string
}

func (d *Dir) MimeFromFilesInDir(dir string) string {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		buf, _ := ioutil.ReadFile(file.Name())
		typeFile := MimeFromFile(buf)
		return typeFile
	}
	return ""
}

func (d *Dir) ReadFilesFromDir(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	fileslist := make([]string, 0)
	for _, file := range files {
		fileslist = append(fileslist, file.Name())
	}
	return fileslist
}

func (d *Dir) NewReadFilesFromDir(dir string) ([]fs.FileInfo, error, *Dir) {
	files, err := ioutil.ReadDir(d.directory)
	return files, err, &Dir{directory: dir}
}

//FileContainsExcel REPLACE THIS METHOD
func (d *Dir) FileContainsExcel(fileName string) string {
	ExcelFiles := []string{"xls", "xlsx", "xlam", ".xlsm", ".xltm", ".xltx", ".xlt", ".xlr"}
	for _, v := range ExcelFiles {
		if strings.Contains(fileName, v) {
			return fileName
		}
	}
	return ""
}
