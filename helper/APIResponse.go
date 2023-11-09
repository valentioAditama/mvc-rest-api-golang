package helper

type APICreatedResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type APIResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}
