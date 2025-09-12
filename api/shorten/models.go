package shorten

type ShortenRequest struct {
	LongURL     string `json:"long_url" validate:"required,url"`
	CustomAlias string `json:"custom_alias" validate:"omitempty,alphanum"`
	Expiry      int    `json:"expiry" validate:"required,min=1"` // in minutes
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
	Expiry   string `json:"expiry"`
}
