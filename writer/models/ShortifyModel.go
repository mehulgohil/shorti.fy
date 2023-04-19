package models

type WriterRequest struct {
	LongURL   string `json:"long_url"`
	UserEmail string `json:"user_email"`
}

type WriterResponse struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
