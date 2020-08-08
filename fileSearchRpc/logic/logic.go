package logic

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/smallnest/rpcx/log"

	"fileSearch/fileSearchRpc/proto"
)

var (
	FilesSyncMap = sync.Map{}
	once         = sync.Once{}
)

func FillFilesMap(path string) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range fileInfos {
		if file.IsDir() {
			FillFilesMap(path + file.Name() + "/")
		} else {
			FilesSyncMap.Store(path+file.Name(), file)
		}
	}
}

func StartSearch(word string, rsp *proto.SearchWordRsp, done chan bool) {
	fileDir := "E:/Work/searchFiles/"
	once.Do(func() {
		FillFilesMap(fileDir)
	})

	searchNames, searchContents := DoSearch(word)
	log.Info(searchNames)
	log.Info(searchContents)
	rsp.Found = true
	rsp.FileNum = 0
	rsp.FileNames = searchNames
	rsp.FileContents = searchContents

	done <- true
}

func DoSearch(word string) (string, string) {
	fileNameChan := make(chan string, 0)
	go SearchFileName(word, fileNameChan)

	fileContentChan := make(chan string, 0)
	go SearchContent(word, fileContentChan)

	fileNameRes := <-fileNameChan
	fileContentRes := <-fileContentChan
	return fileNameRes, fileContentRes
}

func SearchFileName(word string, resChan chan string) {
	resFiles := make([]string, 0)
	FilesSyncMap.Range(func(file, info interface{}) bool {
		filename := file.(string)
		if strings.Contains(filename, word) {
			resFiles = append(resFiles, filename)
			log.Info(filename)
		}
		return true
	})
	resChan <- strings.Join(resFiles, "|")
}

func checkFileExist(filepath string) bool {
	b := true
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			b = false
		}
	}
	return b
}

func SearchContent(word string, resChan chan string) {
	resFiles := make([]string, 0)
	FilesSyncMap.Range(func(filepath, info interface{}) bool {
		if SearchFileContent(word, filepath.(string)) {
			resFiles = append(resFiles, filepath.(string))
		}
		return true
	})
	resChan <- strings.Join(resFiles, "|")
}

func SearchFileContent(word string, filepath string) bool {
	if !checkFileExist(filepath) {
		return false
	}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false
	}
	return strings.Contains(string(bytes), word)
}
