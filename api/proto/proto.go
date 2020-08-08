package proto

type BaseReq struct {
	TraceId string
}

type SearchWordReq struct {
	BaseReq
	Word string
}

type BaseRsp struct {
	ServerId string
}

type SearchWordRsp struct {
	BaseRsp
	Found        bool
	FileNum      int64
	FileNames    string
	FileContents string
	Files        []string
}
