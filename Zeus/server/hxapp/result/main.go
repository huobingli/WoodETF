package result

type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}