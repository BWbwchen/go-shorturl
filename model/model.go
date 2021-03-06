package model

type ShorturlSturct struct {
	Shortname string `json:"shortname"`
	URL       string `json:"url"`
}

func (ShorturlSturct) TableName() string {
	return "shorturl"
}

const (
	Success  = iota
	NotFound = iota
)
