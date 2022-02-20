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

type SearchResult struct {
	FileName string
	LineNo   int64
	Content  string
}

type SearchWordRsp struct {
	BaseRsp
	Found     bool
	FileNum   int64
	SearchRes []SearchResult
}
