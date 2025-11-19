package notify

type NotifyRequest struct {
	Key  string            `json:"key"`
	Data map[string]string `json:"data"`
}
