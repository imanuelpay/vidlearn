package common

type DefaultDataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
