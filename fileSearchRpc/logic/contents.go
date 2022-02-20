package logic

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

func checkFileExist(filepath string) bool {
	b := true
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			b = false
		}
	}
	return b
}

type SearchContentInterface interface {
	SearchContent(word string, filepath string) (found bool, lineno int64, content string)
}

type NormFileSearch struct {
}

func (nfs *NormFileSearch) SearchContent(word string, filepath string) (found bool, lineno int64, content string) {
	found = false
	lineno = 0
	content = ""
	if !checkFileExist(filepath) {
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	enc := mahonia.NewDecoder("GBK")

	for {
		readString, err := reader.ReadString('\n')
		content = enc.ConvertString(readString)
		if err != nil || err == io.EOF {
			break
		}
		if strings.Contains(content, word) {
			found = true
			lineno++
			break
		}
	}
	return
}

type OfficeWordSearch struct {
}

func (ows *OfficeWordSearch) SearchContent(word string, filepath string) (found bool, lineno int64, content string) {
	found = false
	lineno = 0
	content = ""
	return
}

func NewSearchByExt(ext string) SearchContentInterface {
	switch ext {
	case "txt":
		return &NormFileSearch{}
	case "doc", "docx":
		return &OfficeWordSearch{}
	default:
		return &NormFileSearch{}
	}
}
