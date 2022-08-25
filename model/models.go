package model

type Request struct {
	ID      int64               `json:"id"`
	URL     string              `json:"url"`
	Method  string              `json:"method"`
	Headers map[string][]string `json:"headers"`
}

type Response struct {
	ID      string              `json:"id"`
	Status  string              `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int64               `json:"length"`
}
