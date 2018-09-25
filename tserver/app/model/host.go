package model

type HostInfo struct {
	Address string `json:"address"`
	UUID    string `json:"uuid"`
	DNum    int    `json:"unum"`
}

type StartByHostReq struct {
	Rid int64 `json:"rid" form:"rid"`
}
