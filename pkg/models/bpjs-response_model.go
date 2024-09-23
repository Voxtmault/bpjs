package models

type BPJSResponse struct {
	MetaData *MetaData `json:"metaData"`
	Response string    `json:"response"`
}

type MetaData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
