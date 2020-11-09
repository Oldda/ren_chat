package handlers

var(
	NEED_PARAMS_LOST = JsonResponseData{
		Code:10000, //服务器理解客户端的请求，但是拒绝执行请求
		Status:false,
		Msg:"缺少必要参数",
	}
)

type JsonResponseData struct{
	Code int `json:"code"`
	Status bool `json:"status"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func NewJsonResponseData(m JsonResponseData,data ...interface{})(int,JsonResponseData){
	if len(data) > 0{
		m.Data = data[0]
	}
	return 403,m
}