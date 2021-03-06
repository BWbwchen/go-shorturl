package util

import (
	"shorturl/database"
	. "shorturl/model"

	"github.com/google/uuid"
)

type Response struct {
	CheckCode int    `json:"checkcode"`
	ShortName string `json:"short_name"`
}

func GetNewShortName() string {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return id.String()[:7]
}

// CheckValid check whether
// 1. duplicate key
// 2. recursive shortname [NOT YET]
// 3. max short name length is 7
func CheckValid(check string) bool {
	if len(check) > 7 {
		return false
	}
	_, stateCode := database.Find(check)
	return stateCode == NotFound
}

func SendResponse(state int, shortname string) Response {
	return Response{
		CheckCode: state,
		ShortName: shortname,
	}
}
