package logic

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/axgle/mahonia"

	"fileSearch/proto"

	"fileSearch/fileSearchRpc/config"
	"fileSearch/log"
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
			log.GLogger.Infoln(file)
		}
	}
}

func StartSearch(word string, rsp *proto.SearchWordRsp) {
	fileDir := config.GConfig.SearchDir
	once.Do(func() {
		FillFilesMap(fileDir)
	})

	DoSearch(word, &rsp.SearchRes)
	rsp.Found = len(rsp.SearchRes) > 0
	rsp.FileNum = int64(len(rsp.SearchRes))
}

func DoSearch(word string, res *[]proto.SearchResult) {
	fileContentChan := make(chan bool, 0)
	go SearchContent(word, fileContentChan, res)
	<-fileContentChan
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

func SearchContent(word string, doneChan chan bool, res *[]proto.SearchResult) {
	FilesSyncMap.Range(func(file, info interface{}) bool {
		filename := file.(string)
		found, lineno, content := SearchFileContent(word, filename)
		if found {
			log.GLogger.Infof("search %s found %s \n", filename, word)
			*res = append(*res, proto.SearchResult{
				FileName: filename,
				LineNo:   lineno,
				Content:  content,
			})
		}
		return true
	})
	doneChan <- true
}

func SearchFileContent(word string, filepath string) (found bool, lineno int64, content string) {
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

	enc := mahonia.NewEncoder("GBK")

	for {
		readString, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		if strings.Contains(readString, word) {
			found = true
			lineno++
			content = enc.ConvertString(readString)
			break
		}
	}
	return
}
