package models

type WriterRequest struct {
	LongURL string `json:"long_url"`
}

type WriterResponse struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}
