package logic

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"sync"

	"fileSearch/log"
	"fileSearch/proto"

	"fileSearch/fileSearchRpc/config"
	"fileSearch/fileSearchRpc/redis"
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
	if rsp.Found {
		bytes, err := json.Marshal(rsp.SearchRes)
		if err != nil {
			log.GLogger.Errorf("marshal err %v", err)
			return
		}
		redis.RedigoExec("SET", word, string(bytes))
	}
}

func DoSearch(word string, res *[]proto.SearchResult) {
	fileContentChan := make(chan bool, 0)
	go SearchContent(word, fileContentChan, res)
	<-fileContentChan
}

func SearchContent(word string, doneChan chan bool, res *[]proto.SearchResult) {
	FilesSyncMap.Range(func(file, info interface{}) bool {
		filename := file.(string)
		searchInter := NewSearchByExt(path.Ext(filename))
		found, lineno, content := searchInter.SearchContent(word, filename)
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
