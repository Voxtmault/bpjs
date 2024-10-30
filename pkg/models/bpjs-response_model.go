package models

type BPJSResponse struct {
	MetaData *MetaData `json:"metaData"`
	Response string    `json:"response"`
}

type MetaData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type IcareResponse struct {
	MetaData *IcareMetaData `json:"metaData"`
	Response string         `json:"response"`
}

type IcareMetaData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
