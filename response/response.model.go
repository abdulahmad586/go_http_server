package response

import "encoding/json"

type Response struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Err     []byte      `json:"err"`
	Payload interface{} `json:"payload"`
}

func (r *Response) ToJson() []byte {
	js, _ := json.Marshal(r)
	return js
}
