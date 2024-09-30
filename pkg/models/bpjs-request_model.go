package models

// BPJSRequest is used to marshall the many resource editing request struct into a JSON format that is accepted by the BPJS API.
// It's weird, ik dude trust me, but what can i do :D
type BPJSRequest struct {
	Request interface{} `json:"request"`
}
