package dto

type DownloadResponse struct {
	CID  string `json:"cid"`
	Size int    `json:"size"`
	Data []byte `json:"data"`
}
