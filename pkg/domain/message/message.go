package message

type Message struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Status string `json:"status"`
}
