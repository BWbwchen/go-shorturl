package model

import "github.com/dgrijalva/jwt-go"

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

// custom claims
type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}
