package dmcore

type RequestQuery struct {
	Url        string         `json:"url" form:"url"`
	Branch     string         `json:"branch,omitempty" form:"branch"` // branch or tag
	Action     string         `json:"action" form:"action"`           // view|download
	Sub        string         `json:"sub,omitempty" form:"sub"`
	Direction  int            `json:"direction,omitempty" form:"direction"`     // sub dir direction up:1 -> parent dir or down:2 -> sub dir.
	Filename   string         `json:"filename,omitempty" form:"filename"`       // filename -> direct download
	SelectList map[string]int `json:"select_list,omitempty" form:"select_list"` // select file or folders list - max 5
}

type ResponseSuccess struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ResponseError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
