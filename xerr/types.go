package xerr

type BaseMessageResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    any    `json:"Data"`
}
